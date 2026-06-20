package mappers

import (
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func MapCanal(canal string) shared_types.CanalTipo {
	switch canal {
	case "api":
		return shared_types.CanalAPI
	case "web":
		return shared_types.CanalWEB
	case "job":
		return shared_types.CanalJOB
	case "message":
		return shared_types.CanalMSG
	default:
		panic("canal inválido")
	}
}

func MapSistema(sistema string) shared_types.SistemaTipo {
	switch sistema {
	case "bc-legal":
		return shared_types.Legal
	case "bc-documento":
		return shared_types.Documento
	case "bc-profissional":
		return shared_types.Profissional
	case "bc-cliente":
		return shared_types.Cliente
	default:
		panic("sistema inválido")
	}
}
