package adapters

import (
	"time"
	"user-svc/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMetada struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	Exp  int    `json:"exp"`
}

type JwtTokenGenerator struct {
}

func NewJwtTokenGenerator() *JwtTokenGenerator {
	return &JwtTokenGenerator{}
}

func (s *JwtTokenGenerator) Generate(identifier string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = identifier
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return t, nil
}
