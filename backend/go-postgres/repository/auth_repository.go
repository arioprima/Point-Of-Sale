package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/models"
	"log"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, email string) (*models.User, error)
	Register(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
}

type authRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &authRepositoryImpl{DB: db}
}

func (a *authRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email string) (*models.User, error) {
	//TODO implement me
	SQL := `SELECT users.*, roles.name FROM users JOIN roles ON users.role_id = roles.id WHERE email = $1`
	row := tx.QueryRowContext(ctx, SQL, email)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.RoleName, // Add this line to scan the role name
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (a *authRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error) {
	//TODO implement me
	log.Println("Start Register Function")

	// Langkah 1: Check if email already exists
	log.Println("Langkah 1: Memeriksa email yang sudah ada")
	SQL := `SELECT * FROM users WHERE email = $1`
	row := tx.QueryRowContext(ctx, SQL, user.Email)
	var existingUser models.User
	err := row.Scan(
		&existingUser.ID,
		&existingUser.Name,
		&existingUser.Email,
		&existingUser.RoleID,
		&existingUser.Password,
		&existingUser.CreatedAt,
		&existingUser.UpdatedAt,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("Kesalahan memeriksa email yang sudah ada: %v", err)
			return nil, err
		}
	}

	if existingUser.ID != "" {
		log.Println("Email sudah ada")
		return nil, errors.New("email already exists")
	}

	SQL = `INSERT INTO users (id, name, email, password, role_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.ExecContext(
		ctx,
		SQL,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		log.Printf("Kesalahan memasukkan data user: %v", err)
		return nil, err
	}

	log.Println("Register berhasil")
	return user, nil
}
