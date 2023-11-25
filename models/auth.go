package models

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
