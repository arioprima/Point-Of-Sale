package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/initializers"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/models"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/repository"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
	Register(ctx context.Context, request models.RegisterRequest) (models.UserResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: db, Validate: validate}
}

func (auth *AuthServiceImpl) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	// Implementasi goroutine di dalam fungsi Login
	ch := make(chan models.LoginResponse)
	go func() {
		tx, err := auth.DB.Begin()
		if err != nil {
			ch <- models.LoginResponse{}
			return
		}

		defer func() {
			if r := recover(); r != nil {
				err := tx.Rollback()
				if err != nil {
					log.Println("Error rolling back transaction:", err)
				}
			} else {
				err := tx.Commit()
				if err != nil {
					log.Println("Error committing transaction:", err)
				}
			}
		}()

		user, err := auth.AuthRepository.Login(ctx, tx, request.Email)
		if err != nil || user == nil {
			ch <- models.LoginResponse{}
			return
		}

		err = utils.VerifyPassword(user.Password, request.Password)
		if err != nil {
			ch <- models.LoginResponse{}
			return
		}

		config, _ := initializers.LoadConfig(".")

		//generate jwt token
		tokenPayload := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}

		token, err := utils.GenerateToken(config.TokenExpiresIn, tokenPayload, config.TokenSecret)
		if err != nil {
			ch <- models.LoginResponse{}
			return
		}

		ch <- models.LoginResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			RoleID:    user.RoleID,
			RoleName:  user.RoleName,
			TokenType: "Bearer",
			Token:     token,
		}
	}()

	// Menggunakan select untuk menunggu goroutine selesai
	select {
	case response := <-ch:
		return response, nil
	case <-ctx.Done():
		return models.LoginResponse{}, ctx.Err()
	}
}

func (auth *AuthServiceImpl) Register(ctx context.Context, request models.RegisterRequest) (models.UserResponse, error) {
	// Implementasi goroutine di dalam fungsi Register
	ch := make(chan models.UserResponse)
	go func() {
		// Validasi input
		if err := auth.Validate.Struct(request); err != nil {
			ch <- models.UserResponse{}
			return
		}

		// Mulai transaksi
		tx, err := auth.DB.Begin()
		if err != nil {
			ch <- models.UserResponse{}
			return
		}

		// Tunda fungsi untuk menangani rollback atau commit transaksi
		defer func() {
			if r := recover(); r != nil {
				// Terjadi kesalahan, rollback transaksi
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Printf("Kesalahan rollback transaksi: %v", rollbackErr)
				}
				log.Printf("Panic terjadi: %v", r)
			} else {
				// Tidak ada kesalahan, commit transaksi
				if commitErr := tx.Commit(); commitErr != nil {
					log.Printf("Kesalahan commit transaksi: %v", commitErr)
					// Jika terjadi kesalahan commit, rollback transaksi
					if rollbackErr := tx.Rollback(); rollbackErr != nil {
						log.Printf("Kesalahan rollback transaksi setelah kesalahan commit: %v", rollbackErr)
					}
				}
			}
		}()

		// Hash password
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			ch <- models.UserResponse{}
			return
		}

		now := time.Now()
		newUser := models.User{
			ID:        utils.GenerateUUID(),
			Name:      request.Name,
			Email:     request.Email,
			Password:  hashedPassword,
			RoleID:    "3f15230a-2571-41c2-af4c-69136fc0e185",
			CreatedAt: now,
			UpdatedAt: now,
		}

		if newUser.Email == "" {
			ch <- models.UserResponse{}
			return
		}

		user, err := auth.AuthRepository.Register(ctx, tx, &newUser)

		if err != nil {
			// Check if the error is due to an existing email
			if strings.Contains(err.Error(), "email already exists") {
				ch <- models.UserResponse{}
				return
			}
			ch <- models.UserResponse{}
			return
		}
		ch <- models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleName:  user.RoleName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}()

	// Menggunakan select untuk menunggu goroutine selesai
	select {
	case response := <-ch:
		return response, errors.New("email telah digunakan")
	case <-ctx.Done():
		return models.UserResponse{}, ctx.Err()
	}
}
