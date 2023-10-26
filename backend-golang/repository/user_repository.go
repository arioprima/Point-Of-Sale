package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/Point-Of-Sale/models/entity"
	"time"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error)
	Create(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId string) (*entity.User, error)
	FindByUserName(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.User, error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT user_id, firstname, lastname, username, email, password, role, image, created_at, updated_at FROM users WHERE username = ?"
	row := tx.QueryRowContext(
		ctx,
		SQL,
		username,
	)

	var createdTime []uint8
	var updatedTime []uint8

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.UserRole,
		&user.UserImage,
		&createdTime,
		&updatedTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Pengguna tidak ditemukan
			return nil, nil
		}
		return nil, err
	}

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

	return &user, nil
}

func (u *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	//TODO implement me
	checkQuery := "SELECT user_id FROM users WHERE username = ? OR email = ? LIMIT 1"
	var existingUserID string
	err := tx.QueryRowContext(ctx, checkQuery, user.UserName, user.Email).Scan(&existingUserID)

	if err == nil {
		return nil, errors.New("username or email already exists")
	} else if err != sql.ErrNoRows {
		tx.Rollback()
		return nil, err
	}

	SQL := "INSERT INTO users (user_id, firstname, lastname, username, email, password, role, image, created_at, updated_at)" +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err = tx.ExecContext(
		ctx,
		SQL,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Password,
		user.UserRole,
		user.UserImage,
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
	checkQuery := "SELECT user_id FROM users WHERE user_id = ? LIMIT 1"
	var existingUserID string
	err := tx.QueryRowContext(ctx, checkQuery, user.ID).Scan(&existingUserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	SQL := "UPDATE users SET firstname = ?, lastname = ?, image = ?, updated_at = ? WHERE user_id = ?"

	_, err = tx.ExecContext(
		ctx,
		SQL,
		user.FirstName,
		user.LastName,
		user.UserImage,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil

}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error) {
	//TODO implement me
	SQL := "UPDATE users set is_deleted = 1 WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.ID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil

}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT user_id, firstname, lastname, username, email, role, image, created_at, updated_at FROM users WHERE user_id = ?"

	row := tx.QueryRowContext(ctx, SQL, userId)

	var createdTime []uint8
	var updatedTime []uint8

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.UserRole,
		&user.UserImage,
		&createdTime,
		&updatedTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Pengguna tidak ditemukan
			return nil, nil
		}
		return nil, err
	}

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

	return &user, nil
}

func (u *UserRepositoryImpl) FindByUserName(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT user_id, firstname, lastname, username, email, role, password, image, created_at, updated_at FROM users WHERE username = ?"

	row := tx.QueryRowContext(ctx, SQL, username)

	var createdTime []uint8
	var updatedTime []uint8

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.UserRole,
		&user.Password,
		&user.UserImage,
		&createdTime,
		&updatedTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Pengguna tidak ditemukan
			return nil, nil
		}
		return nil, err
	}

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

	return &user, nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*entity.User, error) {
	//TODO implement me
	SQL := "SELECT user_id, firstname, lastname, username, email, role, image, created_at, updated_at FROM users WHERE email = ?"

	row := tx.QueryRowContext(ctx, SQL, email)

	var createdTime []uint8
	var updatedTime []uint8

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.UserRole,
		&user.UserImage,
		&createdTime,
		&updatedTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Pengguna tidak ditemukan
			return nil, nil
		}
		return nil, err
	}

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

	return &user, nil
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.User, error) {
	//TODO implement me
	SQL := "SELECT user_id, firstname, lastname, username, email, role, image, created_at, updated_at FROM users"

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
			&user.UserRole,
			&user.UserImage,
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
