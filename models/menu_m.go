package models

type Menu struct {
	Table     string    `json:"table"`
	AppId     string    `json:"appid" validate:"required"`
	IdCompany string    `json:"idcompany" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	Icon      string    `json:"icon"`
	Desc      string    `json:"desc"`
	Submenu   []Submenu `json:"submenu" validate:"required"`
}
type Submenu struct {
	IdSubmenu int    `json:"idsubmenu" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Url       string `json:"url"`
	Icon      string `json:"icon"`
	Desc      string `json:"desc"`
}
