package user

import (
	"context"
	"go-template/config"
	"go-template/internal/entity"
	"go-template/shared/response"
	"go-template/shared/utils"
)

// Service encapsulates the celeb logic.
type Service struct {
	cfg      config.Config
	repoUser IRepository
}

// NewService creates and returns a new celeb service
func NewService(cfg config.Config, repoUser IRepository) Service {
	return Service{cfg, repoUser}
}

// Query return all users
func (s Service) Query(ctx context.Context, sortBy, viewerID string, offset, limit int) ([]entity.User, int, error) {
	return s.repoUser.Query(ctx, sortBy, offset, limit)
}

// Get returns the user with the specified user ID or username.
func (s Service) Get(ctx context.Context, userID, viewerID string) (entity.User, error) {
	if !utils.IsValidUUID(userID) {
		return s.repoUser.GetByUsername(ctx, userID)
	}
	return s.repoUser.GetByID(ctx, userID, viewerID)
}

// GetByEmail returns the user with the specified user email.
func (s Service) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.repoUser.GetByEmail(ctx, email)
}

// Update updates the user with given ID in the storage.
func (s Service) Update(ctx context.Context, req RequestUpdateUser) error {
	user := entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	rowsAffected, err := s.repoUser.Update(ctx, user)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return response.NotFound()
	}
	return nil
}

// Delete removes the user with given ID from the storage.
func (s Service) Delete(ctx context.Context, id string) error {
	return nil
}
