package postgres

import (
	"context"
	"errors"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type PostRepo struct {
	db *pgx.Conn
}

func NewPostRepo(db *pgx.Conn) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (p *PostRepo) PostByID(ctx context.Context, userID int) (*models.Post, error) {
	return nil, errors.New("Implement me")
}
func (p *PostRepo) SavePost(ctx context.Context, post *models.Post) error {
	return errors.New("Implement me")
}
func (p *PostRepo) UpdatePost(ctx context.Context, post *models.Post) error {
	return errors.New("Implement me")
}
func (p *PostRepo) DeletePostByID(ctx context.Context, id int) error {
	return errors.New("Implement me")
}
func (p *PostRepo) PostsByUserID(ctx context.Context, userID int) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
func (p *PostRepo) PostsByUserEmail(ctx context.Context, email string) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
func (p *PostRepo) PostsByName(ctx context.Context, userID int) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
func (p *PostRepo) PostsByTopicID(ctx context.Context, topicID int) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
func (p *PostRepo) TrendingPosts(ctx context.Context) (*[]models.Post, error) {
	return nil, errors.New("Implement me")
}
