package models 

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}