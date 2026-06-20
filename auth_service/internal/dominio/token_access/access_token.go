package token_access

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"

type AccessToken struct {
	UserID string
	Roles  types.Role
	Autor  types.Autor
	Origem types.Origem
}
