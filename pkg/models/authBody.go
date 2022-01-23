package models 

type AuthenticationRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}