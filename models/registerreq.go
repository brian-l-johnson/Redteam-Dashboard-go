package models

type RegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
