package mappers

import (
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func MapAutorTipo(actorType string) shared_types.AutorTipo {
	switch actorType {
	case "HUMANO":
		return shared_types.AutorHumano
	case "AGENTE":
		return shared_types.AutorAgente
	case "SISTEMA":
		return shared_types.AutorSistema
	default:
		panic("autor tipo inválido")
	}
}
