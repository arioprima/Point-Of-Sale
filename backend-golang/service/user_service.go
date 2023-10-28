package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/models/entity"
	"github.com/arioprima/Point-Of-Sale/models/web/request"
	"github.com/arioprima/Point-Of-Sale/models/web/response"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/arioprima/Point-Of-Sale/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type UserService interface {
	Login(ctx context.Context, request request.UserLoginRequest) (response.LoginResponse, error)
	Create(ctx context.Context, request request.UserCreateRequest) (response.UserResponse, error)
	Update(ctx context.Context, request request.UserUpdateRequest) (response.UserResponse, error)
	Delete(ctx context.Context, userId string) error
	FindById(ctx context.Context, userId string) (response.UserResponse, error)
	FindByUserName(ctx context.Context, username string) (response.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (response.UserResponse, error)
	FindAll(ctx context.Context) ([]response.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: db, Validate: validate}
}

func (service *UserServiceImpl) Login(ctx context.Context, request request.UserLoginRequest) (response.LoginResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return response.LoginResponse{}, err
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

	user, err := service.UserRepository.Login(ctx, tx, request.UserName)
	if err != nil {
		return response.LoginResponse{}, err
	}

	if user == nil {
		return response.LoginResponse{}, errors.New("user not found")
	}

	setConfig, _ := config.LoadConfig(".")

	tokenPayload := map[string]interface{}{
		"user_id":   user.ID,
		"username":  user.UserName,
		"user_role": user.UserRole,
	}

	// Verify password
	verify_error := utils.VerifyPassword(user.Password, request.Password)
	if verify_error != nil {
		return response.LoginResponse{}, verify_error
	}

	// Generate token
	token, err_token := utils.GenerateToken(setConfig.TokenExpiresIn, tokenPayload, setConfig.TokenSecret)

	if err_token != nil {
		return response.LoginResponse{}, err_token
	}

	return response.LoginResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		UserRole:  user.UserRole,
		TokenType: "Bearer",
		Token:     token,
	}, nil
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) (response.UserResponse, error) {
	//TODO implement me
	if err := service.Validate.Struct(request); err != nil {
		return response.UserResponse{}, err
	}

	// Start a transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return response.UserResponse{}, err
	}

	// Defer a function to handle transaction rollback or commit
	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	// Create a user entity
	user := &entity.User{
		ID:        utils.GenerateUUID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  hashedPassword,
		UserRole:  request.UserRole,
		UserImage: request.UserImage,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set a default role if it's not provided
	if user.UserRole == nil {
		role := "employee"
		user.UserRole = &role
	}

	// Check for empty username, email, and password
	if user.UserName == "" {
		return response.UserResponse{}, fmt.Errorf("error: username is empty")
	} else if user.Email == "" {
		return response.UserResponse{}, fmt.Errorf("error: email is empty")
	} else if user.Password == "" {
		return response.UserResponse{}, fmt.Errorf("error: password is empty")
	}

	// Insert the user into the database
	user, err = service.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return response.UserResponse{}, err
	}

	// Return the user response

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		UserRole:  user.UserRole,
		UserImage: user.UserImage,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) (response.UserResponse, error) {
	//TODO implement me
	err := service.Validate.Struct(request)

	if err != nil {
		// bikin json response error
		return response.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				log.Println("Error rolling back transaction:", err)
			}
			panic(err)
		} else {
			err := tx.Commit()
			if err != nil {
				log.Println("Error committing transaction:", err)
			}
		}
	}()

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)

	// bikin kondisi jika user tidak ditemukan maka kmbalikan response error
	if err != nil {
		return response.UserResponse{}, err
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserImage = request.UserImage
	user.UpdatedAt = time.Now()

	user, err = service.UserRepository.Update(ctx, tx, user)

	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		UserRole:  user.UserRole,
		UserImage: user.UserImage,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId string) error {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}

	defer func() {
		err := recover()
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			panic(err)
		} else {
			err := tx.Commit()
			if err != nil {
				return
			}
		}
	}()

	err = service.UserRepository.Delete(ctx, tx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return response.UserResponse{}, err
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

	user, err := service.UserRepository.FindById(ctx, tx, userId)

	if err != nil {
		// Ganti penggunaan panic dengan mengembalikan kesalahan
		return response.UserResponse{}, err
	}

	if user == nil {
		// Pengguna tidak ditemukan
		return response.UserResponse{}, errors.New("user not found")
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		UserRole:  user.UserRole,
		UserImage: user.UserImage,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindByUserName(ctx context.Context, username string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return response.UserResponse{}, err
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

	user, err := service.UserRepository.FindByUserName(ctx, tx, username)

	if err != nil {
		// Ganti penggunaan panic dengan mengembalikan kesalahan
		return response.UserResponse{}, err
	}

	if user == nil {
		// Pengguna tidak ditemukan
		return response.UserResponse{}, errors.New("user not found")
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		UserRole:  user.UserRole,
		UserImage: user.UserImage,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, email string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return response.UserResponse{}, err
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

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)

	if err != nil {
		// Ganti penggunaan panic dengan mengembalikan kesalahan
		return response.UserResponse{}, err
	}

	if user == nil {
		// Pengguna tidak ditemukan
		return response.UserResponse{}, errors.New("user not found")
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		UserRole:  user.UserRole,
		UserImage: user.UserImage,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
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

	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var userResponses []response.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserName:  user.UserName,
			Email:     user.Email,
			UserRole:  user.UserRole,
			UserImage: user.UserImage,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userResponses, nil
}
