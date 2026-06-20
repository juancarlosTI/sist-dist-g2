package identidade

import "context"

type UserRepository interface {
	Salvar(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type ExternalIdentityRepository interface {
	Salvar(ctx context.Context, ei *ExternalIdentity) error
	FindByProviderAndProviderID(ctx context.Context, provider string, providerUserID string) (*ExternalIdentity, error)
	FindByUserID(ctx context.Context, userID string) ([]*ExternalIdentity, error)
	Delete(ctx context.Context, provider string, providerUserID string) error
}
