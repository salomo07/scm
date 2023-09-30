package models

type DefaultResponse struct {
	Status  int    `json:"status"`
	Messege string `json:"messege"`
}

type SessionData struct {
	IdCompany string `json:"idcompany"`
	IdUser    string `json:"iduser"`
	AppId     string `json:"appid"`
	UserCDB   string `json:"ucdb"`
	PassCDB   string `json:"pcdb"`
}
