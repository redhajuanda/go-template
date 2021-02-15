package user

// RequestCreateUser struct
type RequestCreateUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

// RequestUpdateUser represents an user update request.
type RequestUpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// LoginOrRegisterRequest request body
type LoginOrRegisterRequest struct {
	IDToken string `json:"id_token"`
}
