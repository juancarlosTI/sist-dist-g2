package mappers

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"

func RolesTipoToStrings(roles []types.RoleTipo) []string {
	out := make([]string, 0, len(roles))

	for _, r := range roles {
		out = append(out, r.String())
	}

	return out
}

// func StringsToRoles(values string) types.RoleTipo {
// 	// roles := make([]types.RoleTipo, 0, len(values))

// 	// for _, v := range values {
// 	// 	rt, err := types.ParseRoleTipo(v)
// 	// 	if err != nil {
// 	// 		continue // ignora inválido, não quebra tudo
// 	// 	}

// 	// 	roles = append(roles, rt)
// 	// }
// 	strings, err := types.ParseRoleTipo(values)
// 	if err != nil {
// 		return strings
// 	}
// 	return strings
// }

func StringToRole(value string) types.RoleTipo {
	rt, err := types.ParseRoleTipo(value)
	if err != nil {
		return types.RoleProfissional
	}
	return rt
}
