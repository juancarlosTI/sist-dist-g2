package identidade

import (
	"fmt"
	"time"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/common"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/common"
	evento_uuid "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos/uuid"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type User struct {
	id            common.UserID
	email         string
	nome          string
	password_hash *string
	roles         shared_types.Role
	eventos       []EventoAuth
}

func (u *User) registrarEvento(e EventoAuth) error {

	u.eventos = append(u.eventos, e)

	return nil
}

// Limpar eventos

func (u *User) ID() string {
	return u.id.String()
}

func (u *User) HasPassword() bool {
	return u.password_hash != nil
}

func (u *User) CanLoginWithPassword() bool {
	return u.password_hash != nil
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Nome() string {
	return u.nome
}

// func (u *User) Roles() []string {

// 	valores := make([]string, 0, len(u.roles.Tipo))

// 	for _, v := range u.roles.Tipo {
// 		valores = append(valores, v.String())
// 	}

// 	return valores
// }

func (u *User) Role() string {

	return u.roles.Tipo.String()
}

func (u *User) Password() string {
	if u.password_hash == nil {
		return ""
	}
	return *u.password_hash
}

func RegistrarUsuario(
	nome_usuario string,
	email string,
	password string,
	hashSvc service.PasswordService,
	ctx shared_context.EventContext) (*User, error) {
	if email == "" || password == "" {
		return nil, fmt.Errorf("email and password are required")
	}

	id_user, err := common.NovoUserID()
	if err != nil {
		return nil, err
	}

	hash_password, err := hashSvc.Hash(password)
	if err != nil {
		return nil, err
	}

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return nil, err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	// Role padrão
	roles := shared_types.Role{
		Tipo: ctx.Role.Tipo,
	}

	user := User{
		id:            id_user,
		email:         email,
		nome:          nome_usuario,
		password_hash: &hash_password,
		roles:         roles,
		eventos:       []EventoAuth{},
	}

	payload := UsuarioCriadoPayload{
		UserID: user.id,
		Email:  user.email,
		Nome:   user.nome,
		Roles:  user.roles,
	}

	base := common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    string(user.id.String()),
		AgregadoTipo:  "authentication",
		OcorreuAs:     time.Now().UTC(),
		CorrelacaoID:  ctx.CorrelacaoID,
		CausalidadeID: ctx.CausalidadeID,
		Autor: shared_types.Autor{
			Tipo: ctx.Autor.Tipo,
			ID:   ctx.Autor.ID,
		},
		Origem: shared_types.Origem{
			Canal:   ctx.Origem.Canal,
			Sistema: ctx.Origem.Sistema,
		},
		Role: shared_types.Role{
			Tipo: ctx.Role.Tipo,
		},
	}

	hash, err := common.GerarHashEvento(base, payload)
	if err != nil {
		return nil, err
	}

	base.Hash = hash

	evento := UsuarioCriado{
		base:    base,
		payload: payload,
	}

	// EventSourcing é overengeeniring?
	user.registrarEvento(evento)

	return &user, nil
}

func RegistrarUsuarioOIDC(
	email string,
	ctx shared_context.EventContext,
) (*User, error) {

	if email == "" {
		return nil, fmt.Errorf("email is required")
	}

	id_user, err := common.NovoUserID()
	if err != nil {
		return nil, err
	}

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return nil, err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	// Role padrão
	roles := shared_types.Role{
		Tipo: ctx.Role.Tipo,
	}

	nome := "ainda nao definido"

	user := User{
		id:            id_user,
		email:         email,
		nome:          nome,
		password_hash: nil, // 👈 sem senha
		roles:         roles,
		eventos:       []EventoAuth{},
	}

	payload := UsuarioCriadoPayload{
		UserID: id_user,
	}

	base := common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    string(user.id.String()),
		AgregadoTipo:  "authentication",
		OcorreuAs:     time.Now().UTC(),
		CorrelacaoID:  ctx.CorrelacaoID,
		CausalidadeID: ctx.CausalidadeID,
		Autor: shared_types.Autor{
			Tipo: ctx.Autor.Tipo,
			ID:   ctx.Autor.ID,
		},
		Origem: shared_types.Origem{
			Canal:   ctx.Origem.Canal,
			Sistema: ctx.Origem.Sistema,
		},
		Role: shared_types.Role{
			Tipo: ctx.Role.Tipo,
		},
	}

	hash, err := common.GerarHashEvento(base, payload)
	if err != nil {
		return nil, err
	}

	base.Hash = hash

	evento := UsuarioCriado{
		base:    base,
		payload: payload,
	}

	user.registrarEvento(evento)

	return &user, nil
}

func (u *User) AlterarEmail() error {
	return nil
}

func (u *User) AlterarSenha() error {
	// Se não tiver senha, pedir para cadastrar uma.
	return nil
}

func (u *User) Eventos() []EventoAuth {
	return u.eventos
}

func RehidratarUser(
	id string,
	email string,
	nome string,
	password_hash *string,
	roles shared_types.Role,
) *User {

	recriar_agregado, err := common.NovoExternalUserID(id)
	if err != nil {
		return nil
	}

	return &User{
		id:            recriar_agregado,
		email:         email,
		nome:          nome,
		password_hash: password_hash,
		roles:         roles,
		eventos:       []EventoAuth{},
	}
}
