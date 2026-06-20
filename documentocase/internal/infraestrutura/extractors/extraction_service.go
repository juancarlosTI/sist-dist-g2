package extractors

import (
	"context"
	"strings"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type ExtractionService struct {
	extractors []ports.Extractor
}

func NewExtractionService() *ExtractionService {

	return &ExtractionService{
		extractors: []ports.Extractor{
			// NewMuPDFExtractor(),
			NewPDFCPUExtractor(),
			NewPdfGoExtractor(),
		},
	}
}

func (s *ExtractionService) ExtrairTexto(
	ctx context.Context,
	filePath string,
) (*ports.OCRResult, error) {

	for _, extractor := range s.extractors {

		result, err :=
			extractor.Extract(
				ctx,
				filePath,
			)

		if err != nil {
			continue
		}

		if len(strings.TrimSpace(result.Texto)) > 50 {
			return result, nil
		}
	}

	return &ports.OCRResult{
		Texto: "",
	}, nil
}
