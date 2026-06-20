package types

import "fmt"

type RoleTipo string
type AutorTipo string
type CanalTipo string
type SistemaTipo string

// type ProfileType string

func (a AutorTipo) String() string   { return string(a) }
func (a RoleTipo) String() string    { return string(a) }
func (a CanalTipo) String() string   { return string(a) }
func (a SistemaTipo) String() string { return string(a) }

type Role struct {
	Tipo RoleTipo
}

func (r RoleTipo) IsValid() bool {
	switch r {
	case RoleCliente, RoleProfissional, RoleBackoffice:
		return true
	default:
		return false
	}
}

type Autor struct {
	Tipo AutorTipo
	ID   string
}

type Origem struct {
	Canal   CanalTipo
	Sistema SistemaTipo
}

// const (
// 	ProfileProfissional ProfileType = "PROFISSIONAL"
// 	ProfileBackoffice   ProfileType = "BACKOFFICE"
// 	ProfileCliente      ProfileType = "CLIENTE"
// )

const (
	RoleCliente      RoleTipo = "CLIENTE"
	RoleProfissional RoleTipo = "PROFISSIONAL"
	RoleBackoffice   RoleTipo = "BACKOFFICE"
)

const (
	AutorHumano  AutorTipo = "HUMANO"
	AutorAgente  AutorTipo = "AGENTE"
	AutorSistema AutorTipo = "SISTEMA"
)

const (
	CanalAPI CanalTipo = "api"
	CanalWEB CanalTipo = "web"
	CanalJOB CanalTipo = "job"
	CanalMSG CanalTipo = "message"
)

const (
	Frontend     SistemaTipo = "front-end"
	Auth         SistemaTipo = "auth-service"
	Legal        SistemaTipo = "bc-legal"
	Documento    SistemaTipo = "bc-documento"
	Profissional SistemaTipo = "bc-profissional"
	Cliente      SistemaTipo = "bc-cliente"
)

const (
	ErrInvalidRequest     = "invalid_request"
	ErrRegistrationFailed = "registration_failed"
	ErrUnauthorized       = "unauthorized"
)

func ParseSistemaTipo(v string) (SistemaTipo, error) {

	switch SistemaTipo(v) {

	case Legal,
		Documento,
		Profissional,
		Cliente:

		return SistemaTipo(v), nil

	default:
		return "", fmt.Errorf("sistema tipo inválido: %s", v)
	}
}

func ParseCanalTipo(v string) (CanalTipo, error) {

	switch CanalTipo(v) {

	case CanalAPI,
		CanalWEB,
		CanalJOB,
		CanalMSG:

		return CanalTipo(v), nil

	default:
		return CanalAPI, fmt.Errorf("canal tipo inválido: %s", v)
	}
}

func ParseRoleTipo(v string) (RoleTipo, error) {

	switch RoleTipo(v) {

	case
		RoleProfissional,
		RoleBackoffice:
		return RoleTipo(v), nil

	default:
		return "", fmt.Errorf("role tipo inválido: %s", v)
	}
}

func ParseAutorTipo(v string) (AutorTipo, error) {

	switch AutorTipo(v) {

	case AutorHumano,
		AutorAgente,
		AutorSistema:
		return AutorTipo(v), nil

	default:
		return AutorHumano, fmt.Errorf("autor tipo inválido: %s", v)
	}
}
