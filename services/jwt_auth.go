package services

import (
	"clean-go-echo/library"
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTAuth_MethodService interface {
	Authorize(tokenString string) (jwt.Claims, bool, error)
}

type JWTAuthService struct {
	env library.Env
}

func ModuleJwtService(env library.Env) JWTAuth_MethodService {
	return JWTAuthService{
		env: env,
	}
}

func (s JWTAuthService) Authorize(tokenString string) (jwt.Claims, bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.JWTSecret), nil
	})

	if token.Valid {
		return token.Claims.(jwt.MapClaims), true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, false, errors.New("token malformed")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, false, errors.New("token invalid")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, false, errors.New("token expired")
	}

	return nil, false, errors.New("couldn't handle token")
}
