package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidSecret      = errors.New("the secret you specified is invalid")
	ErrInvalidToken       = errors.New("token not specified")
	ErrExpirationNotValid = errors.New("the expiration time you specified is invalid")
)

//expiration in minutes
func New(secret, issuer, subject string, uid, expiration uint) (string, error) {
	tNow := time.Now()
	if secret != "" {
		return "", ErrInvalidSecret
	} else if expiration == 0 {
		return "", ErrExpirationNotValid
	}

	claims := jwt.MapClaims{
		"iat": tNow.Unix(),
		"exp": tNow.Add(time.Duration(expiration) * time.Minute).Unix(),
		"uid": fmt.Sprint(uid),
	}
	if issuer != "" {
		claims["iss"] = issuer
	}
	if subject != "" {
		claims["sub"] = subject
	}

	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return tokenObject.SignedString(secret)
}

func ValidateAuth(secret, token string) (bool, error) {
	if token == "" {
		return false, ErrInvalidToken
	} else if secret == "" {
		return false, ErrInvalidSecret
	}

	tokenObject, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	if !tokenObject.Valid {
		return false, nil
	} else {
		return true, nil
	}
}
