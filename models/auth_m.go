package models

type LoginInput struct {
	IdCompany string `json:"idcompany" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
