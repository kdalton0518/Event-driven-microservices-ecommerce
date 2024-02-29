package encryption

type CustomTokenGenerator struct {
}

func NewCustomTokenGenerator() *CustomTokenGenerator {
	return &CustomTokenGenerator{}
}

func (p *CustomTokenGenerator) Generate(identifier string) (string, error) {
	return identifier, nil
}
