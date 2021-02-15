package entity

import "time"

// UserType custom type
type UserType string

// User represents a user entity.
type User struct {
	ID              string     `json:"id" db:"id"`
	Username        string     `json:"username"`
	FirstName       string     `json:"first_name" db:"first_name"`
	LastName        string     `json:"last_name" db:"last_name"`
	Email           string     `json:"email" db:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" db:"email_verified_at"`
	Password        string     `json:"-" db:"password"`
	SSOSource       string     `json:"sso_source" db:"sso_source"`
	ProfilePic      string     `json:"profile_pic" db:"profile_pic"`
	IsActive        *bool      `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetUsername returns the user username.
func (u User) GetUsername() string {
	return u.Email
}

// GetPassword return user password
func (u User) GetPassword() string {
	return u.Password
}

// GetType returns user type
func (u User) GetType() string {
	return "user"
}
