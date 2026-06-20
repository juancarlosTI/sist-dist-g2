package mappers_claim

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func ToClaims(at *token_access.AccessToken) Claims {

	// tipos := make([]string, 0, len(at.Roles.Tipo))

	// for _, r := range at.Roles.Tipo {
	// 	tipos = append(tipos, r.String())
	// }

	role := at.Roles.Tipo.String()

	return Claims{
		UserID: at.UserID,
		Roles:  role,
		Autor: map[string]string{
			"tipo": at.Autor.Tipo.String(),
			"id":   at.Autor.ID,
		},
		Origem: map[string]string{
			"canal":   at.Origem.Canal.String(),
			"sistema": at.Origem.Sistema.String(),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(15 * time.Minute),
			),
		},
	}
}

func ToDomain(c Claims) token_access.AccessToken {

	// roles := types.Role{
	// 	Tipo: make([]types.RoleTipo, 0),
	// }

	// tipos := c.Roles["tipo"]
	// for _, r := range tipos {
	// 	role := types.RoleTipo(r)

	// 	if !role.IsValid() {
	// 		continue // ou erro
	// 	}

	// 	roles.Tipo = append(roles.Tipo, role)
	// }

	roleTipo, err := types.ParseRoleTipo(c.Roles)
	if err != nil {
		roleTipo = types.RoleProfissional // fallback
	}

	role := types.Role{
		Tipo: roleTipo,
	}

	return token_access.AccessToken{
		UserID: c.UserID,
		Roles:  role,
		Autor: types.Autor{
			Tipo: types.AutorTipo(c.Autor["tipo"]),
			ID:   c.Autor["id"],
		},
		Origem: types.Origem{
			Canal:   types.CanalTipo(c.Origem["canal"]),
			Sistema: types.SistemaTipo(c.Origem["sistema"]),
		},
	}
}
