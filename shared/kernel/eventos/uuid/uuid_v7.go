package uuid

import (
	"github.com/google/uuid"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func NewEventoID() (types.EventoID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return types.EventoID(id.String()), nil
}
