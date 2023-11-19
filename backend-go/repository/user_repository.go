package repository

import (
	"context"
	"database/sql"
	"github.com/arioprima/blog_web/models/entity"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Delete(context.Context, *sql.Tx, *entity.User) error
	FindById(ctx context.Context, tx *sql.Tx, userId string) (*entity.User, error)
	FindByUserName(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.User, error)
}
