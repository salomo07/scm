package models

type LoginInput struct {
	IdCompany string `json:"idcompany" validate:"required"`
	AppId     string `json:"appid" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
