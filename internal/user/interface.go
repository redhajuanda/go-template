package user

import (
	"context"
	"go-template/internal/entity"
)

// IService encapsulates usecase logic for users.
type IService interface {
	// Query return all users
	Query(ctx context.Context, sortBy, viewerID string, offset, limit int) ([]entity.User, int, error)
	// Get returns the user with the specified user ID or username.
	Get(ctx context.Context, userID, viewerID string) (entity.User, error)
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, req RequestUpdateUser) error
	// Delete removes the user with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// IRepository encapsulates the logic to access users from the data source.
type IRepository interface {
	// Query return all users
	Query(ctx context.Context, sortBy string, offset, limit int) ([]entity.User, int, error)
	// GetByID returns the user with the specified user ID.
	GetByID(ctx context.Context, userID, viewerID string) (entity.User, error)
	// GetByEmail returns the user with the specified user email.
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	// GetByUsername returns the user with the specified username.
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	// IsUserExist checks wether user exists
	IsUserExist(ctx context.Context, userID string) (bool, error)
	// IsUserExistByEmail checks wether user exists by email address
	IsUserExistByEmail(ctx context.Context, email string) (bool, error)
	// IsUserExistByUsername checks wether user exists by username
	IsUserExistByUsername(ctx context.Context, username string) (bool, error)
	// Create saves a new user in the storage.
	Create(ctx context.Context, user entity.User) (rowsAffected int, err error)
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, user entity.User) (rowsAffected int, err error)
	// Delete removes the user with given ID from the storage.
	Delete(ctx context.Context, id string) error
}
