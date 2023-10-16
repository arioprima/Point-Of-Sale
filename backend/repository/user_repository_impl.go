package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/blog_web/models/entity"
	"time"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	//TODO implement me
	SQL := "INSERT INTO users (id, firstname, lastname, username, email, password, role, image, created_at, updated_at)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Password,
		user.Role,
		user.Image,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	//TODO implement me
	SQL := "UPDATE users SET firstname = ?," +
		"lastname = ?," +
		"Image = ?," +
		"updated_at = ?" +
		"WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.FirstName,
		user.LastName,
		user.Image,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *entity.User) error {
	//TODO implement me
	SQL := "DELETE FROM users WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username," +
		"email, role, image, created_at, updated_at" +
		"FROM users WHERE id = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		userId,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	user := entity.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return &user, errors.New("user not found")
}

func (u *UserRepositoryImpl) FindByUserName(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username," +
		"email, role, image, created_at, updated_at" +
		"FROM users WHERE username = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		username,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	user := entity.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return &user, errors.New("user not found")
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username," +
		"email, role, image, created_at, updated_at" +
		"FROM users WHERE email = ?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		email,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	user := entity.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return &user, errors.New("user not found")
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.User, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, role, image, created_at, updated_at FROM users"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		user := entity.User{}
		var createdTime []uint8
		var updatedTime []uint8 // Menambahkan variabel untuk kolom updated_at
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.Image,
			&createdTime,
			&updatedTime, // Memindai kolom updated_at ke variabel updatedTime
		)

		if err != nil {
			return nil, err
		}

		// Konversi createdTime dan updatedTime ke tipe data time.Time
		createdTimeStr := string(createdTime)
		parsedCreatedTime, err := time.Parse("2006-01-02 15:04:05", createdTimeStr)
		if err != nil {
			return nil, err
		}
		user.CreatedAt = parsedCreatedTime

		updatedTimeStr := string(updatedTime)
		parsedUpdatedTime, err := time.Parse("2006-01-02 15:04:05", updatedTimeStr)
		if err != nil {
			return nil, err
		}
		user.UpdatedAt = parsedUpdatedTime

		users = append(users, &user)

	}

	return users, nil
}
