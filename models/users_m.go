package models

type User struct {
	IdUser      string      `json:"_id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Username    string      `json:"username" validate:"required"`
	Password    string      `json:"password" validate:"required"`
	IdCompany   string      `json:"table" validate:"required"`
	Role        Role        `json:"role"`
	AccessMenu1 AccessMenu1 `json:"accessmenu1"`
	AccessMenu2 []Contact   `json:"accessmenu2"`
}
type Role struct {
	IdRole string `json:"_id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Desc   string `json:"desc"`
}
type AccessMenu1 struct {
	IdAccessMenu1 string `json:"_id"`
	Idmenu1       string `json:"idmenu1"`
	Create        bool   `json:"create"`
	Read          bool   `json:"read"`
	Update        bool   `json:"update"`
	Delete        bool   `json:"delete"`
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
	Name    string `json:"name"`
	Url     string `json:"url"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
}
type Menu2 struct {
	IdMenu2 string `json:"_id"`
	IdMenu1 string `json:"idmenu1"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
}
