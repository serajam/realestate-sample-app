package e2e

type UserSignUpRequest struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=30"`
	Surname  string `json:"surname" validate:"omitempty,min=2,max=30"`
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type authToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
} // @name AuthToken
