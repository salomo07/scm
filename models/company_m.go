package models

type Company struct {
	IdCompany string `json:"_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Alias     string `json:"alias" validate:"required"`
	Level     string `json:"iduser" validate:"required"`
	Table     string `json:"table" validate:"required"`
}
type CompanyEdit struct {
	IdCompany string `json:"_id" validate:"required"`
	Rev       string `json:"_rev"`
	Name      string `json:"name" validate:"required"`
	Alias     string `json:"alias" validate:"required"`
	Level     string `json:"iduser" validate:"required"`
	Table     string `json:"table" validate:"required"`
}
