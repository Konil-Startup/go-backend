package service

import (
	"context"
	"fmt"

	"github.com/Konil-Startup/go-backend/internal/models"
)

func (s *Service) SaveUser(ctx context.Context, user *models.User) error {
	const op = "service.SaveUser"
	zeroValueUser := models.User{}
	if user == nil || *user == zeroValueUser {
		return ErrEmpty
	}

	if err := s.repo.User.SaveUser(ctx, user); err != nil {
		// TODO
		// if err == errors.ErrUserExists // check for dublicate
		// add metrics
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) UserByID(ctx context.Context, id int) (*models.User, error) {
	const op = "service.UserByID"
	var user *models.User

	user, err := s.repo.UserByID(ctx, id)
	if err != nil {
		// handle domain level errors
		// handle
		// return service layer errors
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

type MetricSuccessfulRequest struct{}
type MetricFailedRequest struct{}

var (
	metricSuccessfulRequest = MetricSuccessfulRequest{}
	metricFailedRequest     = MetricFailedRequest{}
)

func (a *MetricSuccessfulRequest) Inc() {}
func (a *MetricFailedRequest) Inc()     {}

func (s *Service) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	const op = "service.UserByEmail"
	var user *models.User

	// add metrics for successful and unsuccessful execution
	user, err := s.repo.UserByEmail(ctx, email)
	if err != nil {
		// handle repo errs using service layers erorrs
		metricFailedRequest.Inc()
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	metricSuccessfulRequest.Inc()

	return user, nil
}
