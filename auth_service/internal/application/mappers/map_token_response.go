package mappers

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/dto"
)

func MapperTokenResponse(
	accessToken string,
	refreshToken string,
) *dto.TokenResponseDTO {

	return &dto.TokenResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
