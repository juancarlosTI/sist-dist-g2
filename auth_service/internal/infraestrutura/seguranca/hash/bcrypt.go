package hash

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (b *BcryptService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *BcryptService) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
