package adapters

type StubTokenGenerator struct {
}

func NewStubTokenGenerator() *StubTokenGenerator {
	return &StubTokenGenerator{}
}

func (p *StubTokenGenerator) Generate(identifier string) (string, error) {
	return identifier, nil
}
