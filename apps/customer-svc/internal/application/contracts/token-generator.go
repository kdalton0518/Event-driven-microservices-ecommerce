package contracts

type TokenGenerator interface {
	Generate(identifier string) (string, error)
}
