package postgres

import (
	"context"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type TopicRepo struct {
	db *pgx.Conn
}

func NewTopicRepo(db *pgx.Conn) *TopicRepo {
	return &TopicRepo{
		db: db,
	}
}

func (t *TopicRepo) SaveTopic(ctx context.Context, topic *models.Topic) error {
	panic("implement me")
}

func (t *TopicRepo) DeleteTopicByID(ctx context.Context, id int) error {
	panic("implement me")
}

func (t *TopicRepo) UpdateTopicByID(ctx context.Context, id int) error {
	panic("implement me")
}
