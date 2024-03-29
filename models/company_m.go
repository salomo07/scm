package models

type Company struct {
	IdCompany       string    `json:"_id"`
	AppId           string    `json:"appid" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Alias           string    `json:"alias" validate:"required"`
	LevelMembership string    `json:"levelmembership" validate:"required"`
	Table           string    `json:"table"`
	UserCDB         string    `json:"usercdb"`
	PassCDB         string    `json:"passcdb"`
	Contact         []Contact `json:"contact"`
	Role            []Role    `json:"role"`
	Users           []string  `json:"users"`
}
type CompanyUpdate struct {
	IdCompany       string    `json:"_id"`
	Rev             string    `json:"_rev"`
	AppId           string    `json:"appid" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Alias           string    `json:"alias" validate:"required"`
	LevelMembership string    `json:"levelmembership" validate:"required"`
	Table           string    `json:"table"`
	UserCDB         string    `json:"usercdb"`
	PassCDB         string    `json:"passcdb"`
	Contact         []Contact `json:"contact"`
	Role            []Role    `json:"role"`
	Users           []string  `json:"users"`
}
type CompanyEdit struct {
	IdCompany       string    `json:"_id"`
	Rev             string    `json:"_rev"`
	AppId           string    `json:"appid"`
	Name            string    `json:"name"`
	Alias           string    `json:"alias"`
	LevelMembership string    `json:"levelmembership"`
	Table           string    `json:"table"`
	UserCDB         string    `json:"usercdb"`
	PassCDB         string    `json:"passcdb"`
	Contact         []Contact `json:"contact"`
	Users           []string  `json:"users"`
}
type Contact struct {
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Mobile string `json:"mobile"`
}
