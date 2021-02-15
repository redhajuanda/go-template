package user

import (
	"context"
	"go-template/internal/entity"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// Repository encapsulates the logic to access users from the data source.
type Repository struct {
	db *pg.DB
}

// NewRepository creates a new user repository
func NewRepository(db *pg.DB) Repository {
	return Repository{db}
}

// Query return all users
func (r Repository) Query(ctx context.Context, sortBy string, offset, limit int) ([]entity.User, int, error) {
	users := make([]entity.User, 0)
	q := r.db.Model(&users).
		Order("user.created_at DESC")
	totalData, err := q.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, "cannot count users")
	}

	if sortBy == "latest" {
		q.Order("created_at DESC")
	} else if sortBy == "populer" {

	}
	err = q.Offset(offset).Limit(limit).Select()
	if err != nil {
		return nil, 0, errors.Wrap(err, "cannot select users")
	}
	return users, totalData, nil
}

// GetByID returns the user with the specified user ID.
func (r Repository) GetByID(ctx context.Context, userID string, viewerID string) (entity.User, error) {
	user := entity.User{ID: userID}
	q := r.db.Model(&user).Column("user.*")
	err := q.WherePK().Select()
	if err != nil {
		return entity.User{}, errors.Wrap(err, "cannot get user")
	}
	return user, nil
}

// GetByEmail returns the user with the specified user email.
func (r Repository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.Model(&user).Where("email=?", email).Select()
	if err != nil {
		return entity.User{}, errors.Wrap(err, "cannot get user")
	}
	return user, nil
}

// GetByUsername returns the user with the specified username.
func (r Repository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.db.Model(&user).Where("username=?", username).Select()
	if err != nil {
		return entity.User{}, errors.Wrap(err, "cannot get user")
	}
	return user, nil
}

// IsUserExist checks wether user exists
func (r Repository) IsUserExist(ctx context.Context, userID string) (bool, error) {
	user := entity.User{ID: userID}
	exist, err := r.db.Model(&user).WherePK().Exists()
	if err != nil {
		return false, errors.Wrap(err, "cannot check user")
	}
	return exist, nil
}

// IsUserExistByEmail checks wether user exists by email address
func (r Repository) IsUserExistByEmail(ctx context.Context, email string) (bool, error) {
	user := entity.User{}
	exist, err := r.db.Model(&user).Where("email=?", email).Exists()
	if err != nil {
		return false, errors.Wrap(err, "cannot check user")
	}
	return exist, nil
}

// IsUserExistByUsername checks wether user exists by username
func (r Repository) IsUserExistByUsername(ctx context.Context, username string) (bool, error) {
	user := entity.User{}
	exist, err := r.db.Model(&user).Where("username=?", username).Exists()
	if err != nil {
		return false, errors.Wrap(err, "cannot check user")
	}
	return exist, nil
}

// Create saves a new user in the storage.s
func (r Repository) Create(ctx context.Context, user entity.User) (int, error) {
	res, err := r.db.Model(&user).Insert()
	if err != nil {
		return 0, errors.Wrap(err, "cannot create user")
	}
	return res.RowsAffected(), nil
}

// Update updates the user with given ID in the storage.
func (r Repository) Update(ctx context.Context, user entity.User) (int, error) {
	res, err := r.db.Model(&user).UpdateNotZero()
	if err != nil {
		return 0, errors.Wrap(err, "cannot update user")
	}
	return res.RowsAffected(), nil
}

// Delete removes the user with given ID from the storage.
func (r Repository) Delete(ctx context.Context, id string) error {
	user := entity.User{ID: id}
	err := r.db.Model(&user).WherePK().Select()
	if err != nil {
		return errors.Wrap(err, "cannot delete user")
	}
	return nil
}
