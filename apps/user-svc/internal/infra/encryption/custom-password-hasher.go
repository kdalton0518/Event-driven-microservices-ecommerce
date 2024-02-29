package encryption

import (
	"strings"
)

type CustomPasswordHasher struct {
}

func NewCustomPasswordHasher() *CustomPasswordHasher {
	return &CustomPasswordHasher{}
}

func (p *CustomPasswordHasher) Hash(pass string) (string, error) {
	return "hashed:" + pass, nil
}

func (p *CustomPasswordHasher) Compare(passStr, hashedPass string) bool {
	return passStr == strings.Split(hashedPass, ":")[1]
}
