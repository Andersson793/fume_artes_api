package routes

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func JwtValidator(tokenString string) (bool, error) {

	key := os.Getenv("JWT_KEY")

	//validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		return []byte(key), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	switch {

	case errors.Is(err, jwt.ErrTokenMalformed):
		err = fmt.Errorf("A error ocurred: That's not even a token \n")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		err = fmt.Errorf("A error ocurred: Invalid signature \n")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet

		err = fmt.Errorf("A error ocurred: Timing is out \n")
	}

	return token.Valid, err

}
