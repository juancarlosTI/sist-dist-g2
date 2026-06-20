package extractors

import (
	"context"
	"strings"

	"github.com/ledongthuc/pdf"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type PDFCPUExtractor struct{}

func NewPDFCPUExtractor() *PDFCPUExtractor {
	return &PDFCPUExtractor{}
}

func (e *PDFCPUExtractor) Extract(
	ctx context.Context,
	filePath string,
) (*ports.OCRResult, error) {

	f, reader, err := pdf.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var text strings.Builder

	totalPages := reader.NumPage()

	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		page := reader.Page(pageIndex)

		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		text.WriteString(content)
		text.WriteString("\n")
	}

	return &ports.OCRResult{
		Texto: text.String(),
	}, nil
}
