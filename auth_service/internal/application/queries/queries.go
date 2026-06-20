package queries

import (
	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/common"
)

type BuscarProcessoPorIDQuery struct {
	UserID dominio_common.UserID
}
