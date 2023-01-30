package repository

import (
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"

	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUser(ctx *gin.Context, arg dto.CreateUserDTO) (models.User, error)
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (us *UserStorage) CreateUser(ctx *gin.Context, arg dto.CreateUserDTO) (models.User, error) {
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
