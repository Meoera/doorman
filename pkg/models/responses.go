package models

var (
	BadCredentialsResponseBody = map[string]interface{} {
		"message": "Bad Credentials!",
	}
	InternalServerErrorResponseBody = map[string]interface{} {
		"message": "An error occured! Try again later!",
	}
)