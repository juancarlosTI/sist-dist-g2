package token_access

type TokenGenerator interface {
	Generate() (string, error)
}

type TokenHasher interface {
	Hash(value string) string
}
