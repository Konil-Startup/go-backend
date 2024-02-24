package repository

import (
	"context"

	"github.com/Konil-Startup/go-backend/internal/models"
	"github.com/Konil-Startup/go-backend/internal/repository/postgres"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Auth
	User
	Topic
	Post
	Comment
}
type Auth interface {
}

type Topic interface {
	SaveTopic(ctx context.Context, topic *models.Topic) error
	DeleteTopicByID(ctx context.Context, id int) error
	UpdateTopicByID(ctx context.Context, id int) error
}

type User interface {
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	UserByID(ctx context.Context, id int) (*models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
}

type Post interface {
	PostByID(ctx context.Context, userID int) (*models.Post, error)
	SavePost(ctx context.Context, post *models.Post) error
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePostByID(ctx context.Context, id int) error

	PostsByUserID(ctx context.Context, userID int) (*[]models.Post, error)
	PostsByUserEmail(ctx context.Context, email string) (*[]models.Post, error)
	PostsByName(ctx context.Context, userID int) (*[]models.Post, error)

	PostsByTopicID(ctx context.Context, topicID int) (*[]models.Post, error)
	TrendingPosts(ctx context.Context) (*[]models.Post, error)
}

type Comment interface {
	SaveComment(ctx context.Context, comment *models.Comment) error

	// CommentsByPostID(ctx, id) 		 - all comments
	// CommentsByPostID(ctx, id, 10) 	 - first 10 comments
	// CommentsByPostID(ctx, id, 10, 2)  - second 10 comments
	CommentsByPostID(ctx context.Context, postID int, limitAndOffset ...int) (*[]models.Post, error)
	CommentsByUserID(ctx context.Context, userID int, limitAndOffset ...int) (*[]models.Post, error)
}

func NewPostgresRepo(db *pgx.Conn) Repository {
	return Repository{
		User:    postgres.NewUserRepo(db),
		Topic:   postgres.NewTopicRepo(db),
		Post:    postgres.NewPostRepo(db),
		Comment: postgres.NewCommentRepo(db),
	}
}
