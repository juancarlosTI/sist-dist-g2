package extractors

import (
	"bytes"
	"context"
	"fmt"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	"github.com/ledongthuc/pdf"
)

type PdfGoExtractor struct{}

func NewPdfGoExtractor() *PdfGoExtractor {
	return &PdfGoExtractor{}
}

func (e *PdfGoExtractor) Extract(
	ctx context.Context,
	filePath string,
) (*ports.OCRResult, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	pdf.DebugOn = true

	f, r, err := pdf.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		panic(err)
	}
	buf.ReadFrom(b)
	content := buf.String()
	fmt.Println(content)

	return &ports.OCRResult{
		Texto: content,
	}, nil
}
