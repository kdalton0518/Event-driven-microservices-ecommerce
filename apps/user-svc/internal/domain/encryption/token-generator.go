package encryption

type TokenGenerator interface {
	Generate(identifier string) (string, error)
}
