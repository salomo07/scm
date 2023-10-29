package models

type User struct {
	IdUser      string        `json:"_id"`
	Nik         string        `json:"nik" validate:"required"`
	Name        string        `json:"name" validate:"required"`
	Username    string        `json:"username" validate:"required"`
	Password    string        `json:"password" validate:"required"`
	Table       string        `json:"table"`
	IdCompany   string        `json:"idcompany"`
	IdRole      string        `json:"idrole" validate:"required"`
	AccessMenu1 []AccessMenu1 `json:"accessmenu1"`
	AccessMenu2 []AccessMenu2 `json:"accessmenu2"`
}
type Role struct {
	Table     string `json:"table"`
	IdCompany string `json:"idcompany" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Desc      string `json:"desc"`
}
type AccessMenu1 struct {
	IdAccessMenu1 string `json:"_id"`
	Idmenu1       string `json:"idmenu1" validate:"required"`
	Create        bool   `json:"create" validate:"required"`
	Read          bool   `json:"read" validate:"required"`
	Update        bool   `json:"update" validate:"required"`
	Delete        bool   `json:"delete" validate:"required"`
}
type AccessMenu2 struct {
	IdAccessMenu2 string `json:"_id"`
	Idmenu2       string `json:"idmenu2"`
	Create        bool   `json:"create"`
	Read          bool   `json:"read"`
	Update        bool   `json:"update"`
	Delete        bool   `json:"delete"`
}
type Menu1 struct {
	IdMenu1 string `json:"_id"`
	Name    string `json:"name" validate:"required"`
	Url     string `json:"url"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
}
type Menu2 struct {
	IdMenu2 string `json:"_id"`
	IdMenu1 string `json:"idmenu1" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Url     string `json:"url"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
}
