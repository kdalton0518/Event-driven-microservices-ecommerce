package encryption

import (
	"strings"
)

type StubPasswordHasher struct {
}

func NewStubPasswordHasher() *StubPasswordHasher {
	return &StubPasswordHasher{}
}

func (p *StubPasswordHasher) Hash(pass string) (string, error) {
	return "hashed:" + pass, nil
}

func (p *StubPasswordHasher) Compare(passStr, hashedPass string) bool {
	return passStr == strings.Split(hashedPass, ":")[1]
}
