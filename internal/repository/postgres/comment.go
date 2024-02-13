package postgres

import (
	"context"
	"errors"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type CommnetRepo struct {
	db *pgx.Conn
}

func NewCommentRepo(db *pgx.Conn) *CommnetRepo {
	return &CommnetRepo{
		db: db,
	}
}
func (c *CommnetRepo) SaveComment(ctx context.Context, comment *models.Comment) error {
	return errors.New("Implement me")
}
func (c *CommnetRepo) CommentsByPostID(ctx context.Context, postID int, limitAndOffset ...int) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
func (c *CommnetRepo) CommentsByUserID(ctx context.Context, userID int, limitAndOffset ...int) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
