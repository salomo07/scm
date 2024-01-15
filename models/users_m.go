package models

type User struct {
	Id        string  `json:"_id" validate:"required"`
	Nik       string  `json:"nik" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Nickname  string  `json:"nickname"`
	Username  string  `json:"username" validate:"required"`
	Password  string  `json:"password" validate:"required"`
	Table     string  `json:"table"`
	IdCompany string  `json:"idcompany" validate:"required"`
	IdRole    string  `json:"idrole" validate:"required"`
	Contact   Contact `json:"contact" validate:"required"`
}

type Role struct {
	Table     string `json:"table"`
	IdCompany string `json:"idcompany" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Desc      string `json:"desc"`
}
type AccessMenu struct {
	Table         string          `json:"table"`
	IdCompany     string          `json:"idcompany" validate:"required"`
	IdRole        string          `json:"idrole" validate:"required"`
	Idmenu        string          `json:"idmenu" validate:"required"`
	AccessSubmenu []AccessSubmenu `json:"accesssubmenu" validate:"required"`
	Create        bool            `json:"create" validate:"required"`
	Read          bool            `json:"read" validate:"required"`
	Update        bool            `json:"update" validate:"required"`
	Delete        bool            `json:"delete" validate:"required"`
}
type AccessMenuUpdate struct {
	IdAccess      string          `json:"_id"`
	Rev           string          `json:"_rev"`
	Table         string          `json:"table"`
	IdCompany     string          `json:"idcompany" validate:"required"`
	IdRole        string          `json:"idrole" validate:"required"`
	Idmenu        string          `json:"idmenu" validate:"required"`
	AccessSubmenu []AccessSubmenu `json:"accesssubmenu" validate:"required"`
	Create        bool            `json:"create" validate:"required"`
	Read          bool            `json:"read" validate:"required"`
	Update        bool            `json:"update" validate:"required"`
	Delete        bool            `json:"delete" validate:"required"`
}
type AccessSubmenu struct {
	Idsubmenu int  `json:"idsubmenu"`
	Create    bool `json:"create"`
	Read      bool `json:"read"`
	Update    bool `json:"update"`
	Delete    bool `json:"delete"`
}
type Menu struct {
	Table     string    `json:"table"`
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

type PublishRedis struct {
	IdCompany string `json:"idcompany" validate:"required"`
	Data      any    `json:"data" validate:"required"`
}
