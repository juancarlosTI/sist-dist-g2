package auth

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"

type AuthContext struct {
	Autor  types.Autor
	Roles  types.Role
	Origem types.Origem
}
