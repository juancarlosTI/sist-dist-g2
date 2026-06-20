package services

import (
	"context"
	"log"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type DocumentoProcessor interface {
	ProcessarDocumento(
		ctx context.Context,
		documentoID string,
		arquivoID string,
	) error
}

type DocumentoProcessorImpl struct {
	storage ports.Storage
	ocr     ports.OCR
}

func NewDocumentoProcessor(
	storage ports.Storage,
	ocr ports.OCR,
) *DocumentoProcessorImpl {

	return &DocumentoProcessorImpl{
		storage: storage,
		ocr:     ocr,
	}
}

func (p *DocumentoProcessorImpl) ProcessarDocumento(
	ctx context.Context,
	documentoID string,
	arquivoID string,
) error {

	log.Printf(
		"processando documento=%s arquivo=%s",
		documentoID,
		arquivoID,
	)

	arquivo, err :=
		p.storage.Download(
			ctx,
			arquivoID,
		)

	if err != nil {
		return err
	}

	defer arquivo.Close()

	resultadoOCR, err :=
		p.ocr.ExtrairTexto(
			ctx,
			arquivo,
		)

	if err != nil {
		return err
	}

	log.Printf(
		"OCR concluído documento=%s texto=%s",
		documentoID,
		resultadoOCR.Texto,
	)

	return nil
}
