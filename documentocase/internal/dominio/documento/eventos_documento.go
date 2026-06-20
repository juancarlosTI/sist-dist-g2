package documento

import (
	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
)

type EventoDocumento interface {
	Nome() string
	Base() dominio_common.EventoBase
	Payload() any
}

type DocumentoCriadoPayload struct {
	DocumentoID dominio_common.DocumentoID `json:"documento_id"`
	ArquivoID   string                     `json:"arquivo_id"`
}

type DocumentoCriado struct {
	base    dominio_common.EventoBase
	payload DocumentoCriadoPayload
}

func (e DocumentoCriado) Nome() string {
	return "DocumentoCriado"
}

func (e DocumentoCriado) Base() dominio_common.EventoBase {
	return e.base
}

func (e DocumentoCriado) Payload() any {
	return e.payload
}

type DocumentoAssociadoAoProcesssoPayload struct {
	DocumentoID dominio_common.DocumentoID        `json:"documento_id"`
	ProcessoID  dominio_common.ExternalProcessoID `json:"processo_id"`
}

type DocumentoAssociadoAoProcesso struct {
	base    dominio_common.EventoBase
	payload DocumentoAssociadoAoProcesssoPayload
}

func (e DocumentoAssociadoAoProcesso) Nome() string {
	return "DocumentoAssociadoAoProcesso"
}

func (e DocumentoAssociadoAoProcesso) Base() dominio_common.EventoBase {
	return e.base
}

func (e DocumentoAssociadoAoProcesso) Payload() any {
	return e.payload
}

type DocumentoAtualizadoPayload struct {
	DocumentoID dominio_common.DocumentoID `json:"documento_id"`
}

type DocumentoAtualizado struct {
	base    dominio_common.EventoBase
	payload DocumentoAtualizadoPayload
}

func (e DocumentoAtualizado) Nome() string {
	return "DocumentoAtualizado"
}

func (e DocumentoAtualizado) Base() dominio_common.EventoBase {
	return e.base
}

func (e DocumentoAtualizado) Payload() any {
	return e.payload
}

type DocumentoInvalidadoPayload struct {
	DocumentoID dominio_common.DocumentoID `json:"documento_id"`
}

type DocumentoInvalidado struct {
	base    dominio_common.EventoBase
	payload DocumentoInvalidadoPayload
}

func (e DocumentoInvalidado) Nome() string {
	return "DocumentoInvalidado"
}

func (e DocumentoInvalidado) Base() dominio_common.EventoBase {
	return e.base
}

func (e DocumentoInvalidado) Payload() any {
	return e.payload
}

type DocumentoArquivadoPayload struct {
	DocumentoID dominio_common.DocumentoID `json:"documento_id"`
}

type DocumentoArquivado struct {
	base    dominio_common.EventoBase
	payload DocumentoArquivadoPayload
}

func (e DocumentoArquivado) Nome() string {
	return "DocumentoArquivado"
}

func (e DocumentoArquivado) Base() dominio_common.EventoBase {
	return e.base
}

func (e DocumentoArquivado) Payload() any {
	return e.payload
}
