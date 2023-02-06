package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type User interface {
	CreateUser(ctx context.Context, arg dto.CreateUserRequestDTO) (models.User, error)
	GetUser(ctx context.Context, username string) (models.User, error)
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (us *UserStorage) CreateUser(ctx context.Context, arg dto.CreateUserRequestDTO) (models.User, error) {
	var user models.User

	query := `INSERT INTO users (
		username,
		hashed_password,
		full_name,
		email
	  ) VALUES (
		$1, $2, $3, $4
	  ) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
	  `

	row := us.db.QueryRowContext(ctx, query, arg.Username, arg.Password, arg.FullName, arg.Email)

	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
	)

	return user, err
}

func (us *UserStorage) GetUser(ctx context.Context, username string) (dto.LoginUserResponseDTO, error) {
	var user (dto.LoginUserResponseDTO, error)

	query := `SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users WHERE username = $1 LIMIT 1`

	row := us.db.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
	)

	return user, err
}
