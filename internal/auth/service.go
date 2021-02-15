package auth

import (
	"context"
	"fmt"
	"go-template/config"
	"go-template/internal/user"
	"go-template/shared/password"
	"go-template/shared/response"
	"go-template/shared/utils"
	"go-template/shared/validator"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Service encapsulates the authentication logic.
type Service struct {
	cfg      config.Config
	repoUser user.IRepository
}

// NewService creates and returns a new auth service
func NewService(cfg config.Config, repoUser user.IRepository) Service {
	return Service{cfg, repoUser}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	err := validator.Validate(req)
	if err != nil {
		return LoginResponse{}, err
	}

	if identity := s.authenticate(ctx, req.Email, req.Password); identity != nil {
		accessToken, expiresAt, refreshToken, err := s.generateJWT(identity)
		return LoginResponse{
			AccessToken:  accessToken,
			ExpiresAt:    expiresAt,
			RefreshToken: refreshToken,
		}, err
	}
	return LoginResponse{}, response.Unauthorized("Invalid email or password")
}

// RefreshToken refresh the access token
func (s Service) RefreshToken(ctx context.Context, req RefreshTokenRequest) (RefreshTokenResponse, error) {
	token, err := utils.VerifyToken(req.RefreshToken, s.cfg.JWT.SigningKey)
	if err != nil {
		return RefreshTokenResponse{}, response.Forbidden("Refresh token is invalid")
	}
	claims := token.Claims.(jwt.MapClaims)
	var tokenType string
	if val, ok := claims["token_type"].(string); ok {
		tokenType = val
	}

	if tokenType != "refresh" {
		return RefreshTokenResponse{}, response.Forbidden("Refresh token is invalid")
	}

	var id string
	if val, ok := claims["id"].(string); ok {
		id = val
	}

	user, err := s.repoUser.GetByID(ctx, id, "")
	if err != nil {
		return RefreshTokenResponse{}, response.Forbidden()
	}

	accessToken, expiresAt, err := s.generateAccessToken(user)
	if err != nil {
		return RefreshTokenResponse{}, errors.Wrap(fmt.Errorf("wow"), "cannot generate token")
	}
	return RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
}

// authenticate authenticates a user using email and password.
// If email and password are correct, an identity is returned. Otherwise, nil is returned.
func (s Service) authenticate(ctx context.Context, email, plainPwd string) Identity {
	// logger.WithParam("user", email)
	user := s.getUser(ctx, email)
	if user == nil {
		return nil
	}
	spew.Dump(user)
	// fmt.Println(plainPwd)
	// fmt.Println(password.HashAndSalt([]byte(plainPwd)))
	fmt.Println(user.GetPassword(), plainPwd)
	fmt.Println(password.ComparePasswords(user.GetPassword(), []byte(plainPwd)))
	if email == user.GetUsername() && password.ComparePasswords(user.GetPassword(), []byte(plainPwd)) {
		// logger.Infof("authentication successful")
		return user
	}

	// logger.Infof("authentication failed")
	return nil
}

func (s Service) getUser(ctx context.Context, email string) Identity {
	user, err := s.repoUser.GetByEmail(ctx, email)
	if err != nil {
		return nil
	}
	return user
}

// generateJWT generates a JWT that encodes an identity.
func (s Service) generateJWT(identity Identity) (accessToken string, expiresAt int64, refreshToken string, err error) {
	accessToken, expiresAt, err = s.generateAccessToken(identity)

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         identity.GetID(),
		"exp":        time.Now().AddDate(1000, 0, 0).Unix(),
		"token_type": "refresh",
	}).SignedString([]byte(s.cfg.JWT.SigningKey))
	return
}

func (s Service) generateAccessToken(identity Identity) (accessToken string, expiresAt int64, err error) {
	// expiresAt = time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix()
	expiresAt = time.Now().Add(1000 * time.Minute).Unix()
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         identity.GetID(),
		"username":   identity.GetUsername(),
		"user_type":  identity.GetType(),
		"exp":        expiresAt,
		"token_type": "access",
	}).SignedString([]byte(s.cfg.JWT.SigningKey))
	return
}
