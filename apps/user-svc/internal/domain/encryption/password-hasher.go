package encryption

type PasswordHasher interface {
	Hash(string) (string, error)
	Comapre(passStr, hashedPass string) (bool, error)
}
