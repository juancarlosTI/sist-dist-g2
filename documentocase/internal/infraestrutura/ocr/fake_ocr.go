package ocr

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type FakeOCR struct{}

func NewFakeOCR() *FakeOCR {
	return &FakeOCR{}
}

func (o *FakeOCR) ExtrairTexto(
	ctx context.Context,
	arquivoID string,
) (*ports.OCRResult, error) {

	return &ports.OCRResult{
		Texto: `
PROCESSO Nº 1234567-89.2026.8.27.0001

AUTOR: João da Silva

RÉU: Empresa XPTO LTDA

Objeto: Ação de Cobrança

Valor da causa: R$ 15.000,00
`,
	}, nil
}
