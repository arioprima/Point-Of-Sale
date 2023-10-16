package service

import (
	"context"
	"database/sql"
	"github.com/arioprima/blog_web/models/entity"
	"github.com/arioprima/blog_web/models/web/request"
	"github.com/arioprima/blog_web/models/web/response"
	"github.com/arioprima/blog_web/repository"
	"github.com/arioprima/blog_web/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: db, Validate: validate}
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) (response.UserResponse, error) {
	//TODO implement me
	err := service.Validate.Struct(request)

	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		ID:        utils.GenerateUUID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  hashedPassword,
		Role:      request.Role,
		Image:     request.Image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = service.UserRepository.Create(ctx, tx, user)

	if err != nil {
		panic(err)
	}
	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) (response.UserResponse, error) {
	//TODO implement me
	err := service.Validate.Struct(request)

	if err != nil {
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)

	if err != nil {
		panic(err)
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Image = request.Image
	user.UpdatedAt = time.Now()

	user, err = service.UserRepository.Update(ctx, tx, user)

	if err != nil {
		panic(err)
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId string) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	user, err := service.UserRepository.FindById(ctx, tx, userId)

	if err != nil {
		panic(err)
	}

	err = service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	user, err := service.UserRepository.FindById(ctx, tx, userId)

	if err != nil {
		panic(err)
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindByUserName(ctx context.Context, username string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	user, err := service.UserRepository.FindByUserName(ctx, tx, username)

	if err != nil {
		panic(err)
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, email string) (response.UserResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
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

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)

	if err != nil {
		panic(err)
	}

	return response.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
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
			Role:      user.Role,
			Image:     user.Image,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return userResponses, nil
}
