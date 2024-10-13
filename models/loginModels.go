package models

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum" example:"VIPIN"`
	Password string `json:"password" validate:"required,min=8" example:"Vipin@123"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
