package models 

type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}