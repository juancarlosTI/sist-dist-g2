package auth

import (
	"fmt"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func (c AuthContext) Role(role types.Role) bool {
	switch role.Tipo {
	case types.RoleProfissional:
		fmt.Println("yo")
		return true

	case types.RoleCliente:
		fmt.Println("yo")
		return true

	case types.RoleBackoffice:
		fmt.Println("yo")
		return true

	default:
		return false
	}
}

func (c AuthContext) Ator(ator types.Autor) bool {
	switch ator.Tipo {
	case types.AutorHumano:
		fmt.Println("yo")
		return true

	case types.AutorAgente:
		fmt.Println("yo")
		return true

	case types.AutorSistema:
		fmt.Println("yo")
		return true

	default:
		return false
	}
}
