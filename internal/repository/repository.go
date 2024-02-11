package repository

import (
	"context"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/Konil-Startup/go-backend/internal/repository/postgres"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	User
}

type User interface {
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	UserByID(ctx context.Context, id int) (*models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
}

func NewPostgresRepo(db *pgx.Conn) Repository {
	return Repository{
		User: postgres.NewUserRepo(db),
	}
}
