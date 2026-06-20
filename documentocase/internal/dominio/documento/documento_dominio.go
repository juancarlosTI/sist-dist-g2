package documento

import (
	"log"
	"time"

	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/common"
	evento_uuid "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos/uuid"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type Documento struct {
	id           dominio_common.DocumentoID
	estado       EstadoDocumento
	versao       int
	origem       shared_types.Origem
	autor        shared_types.Autor
	processosIDs []*dominio_common.ExternalProcessoID
	eventos      []EventoDocumento
}

func (p *Documento) registrarEvento(e EventoDocumento) error {

	err := p.aplicarEvento(e)
	if err != nil {
		return err
	}

	p.eventos = append(p.eventos, e)

	p.versao++

	return nil
}

func (d *Documento) aplicarEvento(e EventoDocumento) error {

	switch e := e.(type) {

	case DocumentoCriado:
		payload := e.Payload()
		log.Println("Payload: ", payload)
		id, err := dominio_common.NovoExternalDocumentoID(d.ID())
		if err != nil {
			return err
		}
		d.id = id

	}
	return nil

}

func (d *Documento) Estado() EstadoDocumento {
	return d.estado
}

func (d *Documento) Versao() int {
	return d.versao
}

func (d *Documento) Processos() []string {

	result := make([]string, len(d.processosIDs))

	for i, id := range d.processosIDs {
		result[i] = *id.String()
	}

	return result
}

func (d *Documento) OrigemCanal() string {
	return d.origem.Canal.String()
}

func (d *Documento) OrigemSistema() string {
	return d.origem.Sistema.String()
}

func (d *Documento) AutorTipo() string {
	return d.autor.Tipo.String()
}

func (d *Documento) AutorID() string {
	return d.autor.ID
}

func (d *Documento) Eventos() []EventoDocumento {
	return d.eventos
}

func (d *Documento) LimparEventos() {
	d.eventos = nil
}

func (d *Documento) GerarSnapshot() DocumentoSnapshot {

	processosIDs := make([]string, len(d.processosIDs))

	for i, processoID := range d.processosIDs {

		processosIDs[i] = *processoID.String()
	}

	return DocumentoSnapshot{
		ID:            d.ID(),
		Estado:        int(d.estado),
		OrigemCanal:   d.origem.Canal.String(),
		OrigemSistema: d.origem.Sistema.String(),
		AutorTipo:     d.autor.Tipo.String(),
		AutorID:       d.autor.ID,
		Versao:        d.versao,
		ProcessosIDs:  processosIDs,
	}
}

// Factory
func CriarDocumento(ctx shared_context.EventContext, arquivo_id string) (*Documento, error) {
	doc_id, err := dominio_common.NovoDocumentoID()
	if err != nil {
		return nil, nil
	}

	d := &Documento{
		id:     doc_id,
		estado: Ativo,
		autor: shared_types.Autor{
			Tipo: ctx.Autor.Tipo,
			ID:   ctx.Autor.ID,
		},
		origem: shared_types.Origem{
			Canal:   ctx.Origem.Canal,
			Sistema: ctx.Origem.Sistema,
		},
		processosIDs: nil,
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

	payload := DocumentoCriadoPayload{
		DocumentoID: doc_id,
		ArquivoID:   arquivo_id,
	}

	base := dominio_common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    d.ID(),
		AgregadoTipo:  "documento",
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

	hash, err := dominio_common.GerarHashEvento(base, payload)
	if err != nil {
		return nil, err
	}

	base.Hash = hash
	evento := DocumentoCriado{
		base:    base,
		payload: payload,
	}

	d.registrarEvento(evento)

	return d, nil
}

func (d *Documento) AssociarDocumentoAoProcesso(processo_id dominio_common.ExternalProcessoID,
	ctx shared_context.EventContext) error {
	if d.estado == Arquivado {
		return nil
	}

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	payload := DocumentoAssociadoAoProcesssoPayload{
		DocumentoID: d.id,
		ProcessoID:  processo_id,
	}
	base := dominio_common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    d.ID(),
		AgregadoTipo:  "documento",
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
	}

	hash, err := dominio_common.GerarHashEvento(base, payload)
	if err != nil {
		return err
	}

	base.Hash = hash
	evento := DocumentoAssociadoAoProcesso{
		base:    base,
		payload: payload,
	}

	d.registrarEvento(evento)
	return nil
}

func (d *Documento) AtualizarConteudo(ctx shared_context.EventContext) error {
	if d.estado == Arquivado {
		return nil
	}

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	payload := DocumentoAtualizadoPayload{
		DocumentoID: d.id,
	}
	base := dominio_common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    d.ID(),
		AgregadoTipo:  "documento",
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
	}

	hash, err := dominio_common.GerarHashEvento(base, payload)
	if err != nil {
		return err
	}

	base.Hash = hash
	evento := DocumentoAtualizado{
		base:    base,
		payload: payload,
	}

	d.registrarEvento(evento)

	return nil
}

func (d *Documento) InvalidarDocumento(motivo string, ctx shared_context.EventContext) error {
	if d.estado == Arquivado {
		return nil
	}

	d.estado = Invalidado

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	payload := DocumentoAtualizadoPayload{
		DocumentoID: d.id,
	}
	base := dominio_common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    d.ID(),
		AgregadoTipo:  "documento",
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
	}

	hash, err := dominio_common.GerarHashEvento(base, payload)
	if err != nil {
		return err
	}

	base.Hash = hash
	evento := DocumentoAtualizado{
		base:    base,
		payload: payload,
	}

	d.registrarEvento(evento)

	return nil
}

func (d *Documento) ArquivarDocumento(ctx shared_context.EventContext) error {
	if d.estado == Arquivado {
		return nil
	}

	d.estado = Arquivado

	evento_id, err := evento_uuid.NewEventoID()
	if err != nil {
		return err
	}

	if ctx.CorrelacaoID == "" {
		ctx.CorrelacaoID = string(evento_id)
	}

	if ctx.CausalidadeID == "" {
		ctx.CausalidadeID = ctx.CorrelacaoID
	}

	payload := DocumentoArquivadoPayload{
		DocumentoID: d.id,
	}
	base := dominio_common.EventoBase{
		EventoID:      evento_id,
		AgregadoID:    d.ID(),
		AgregadoTipo:  "documento",
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
	}

	hash, err := dominio_common.GerarHashEvento(base, payload)
	if err != nil {
		return err
	}

	base.Hash = hash
	evento := DocumentoArquivado{
		base:    base,
		payload: payload,
	}

	d.registrarEvento(evento)

	return nil
}

func (d *Documento) EstaArquivado() bool {
	return d.estado == Arquivado
}

func (d *Documento) EstaInvalidado() bool {
	return d.estado == Invalidado
}

func (d *Documento) ID() string {
	return d.id.String()
}

func ReconstruirAgregado(
	snap *DocumentoSnapshot,
	eventos []EventoDocumento,
) (*Documento, error) {

	var p *Documento

	// if snap != nil {

	// 	pid, err := dominio_common.NovoExternalDocumentoID(snap.ID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	p = &Documento{
	// 		id:     pid,
	// 		estado: EstadoDocumento(snap.Estado),
	// 		versao: snap.Versao,
	// 	}

	// 	// Advogados
	// 	for _, proc_id := range snap.ProcessosIDs {

	// 		aid, err := dominio_common.NewExternalProfissionalID(advID)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 		p.advogadosIDs = append(p.advogadosIDs, aid)
	// 	}

	// } else {
	// 	if len(eventos) == 0 {
	// 		return nil, nil
	// 	}

	// 	pid, err := dominio_common.NovoProcessoExternoID(eventos[0].Base().AgregadoID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	p = &Documento{
	// 		id: pid,
	// 	}
	// }

	for _, evt := range eventos {

		err := p.aplicarEvento(evt)
		if err != nil {
			return nil, err
		}

	}

	if snap != nil && len(eventos) > 0 {

		expected := snap.Versao + 1

		first := eventos[0].Base().Versao

		if first != expected {

			return nil, nil
		}
	}

	if len(eventos) > 0 {

		expected := eventos[0].Base().Versao

		for _, evt := range eventos {

			if evt.Base().Versao != expected {

				return nil, dominio_common.ErrEventStreamCorrompido
			}

			expected++
		}
	} else if snap != nil {

		p.versao = snap.Versao

	}

	p.eventos = nil

	return p, nil
}
