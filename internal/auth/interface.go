package auth

import (
	"context"
)

// IService encapsulates the authentication logic.
type IService interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest) (RefreshTokenResponse, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetUsername() string
	// GetPassword return user password
	GetPassword() string
	// GetType returns user type
	GetType() string
}
