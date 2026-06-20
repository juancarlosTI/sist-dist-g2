package mappers

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/dto"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
)

func ToUserResponseDTO(user *identidade.User) *dto.UserResponseDTO {
	if user == nil {
		return nil
	}

	// roles := make([]string, 0, len(user.Roles()))
	// for _, r := range user.Roles() {
	// 	roles = append(roles, r)
	// }
	role := user.Role()
	notNullName := "NotNull"

	return &dto.UserResponseDTO{
		ID:    user.ID(),
		Email: user.Email(),
		Nome:  &notNullName,
		Roles: role,
	}
}
