package ocr

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/extractors"
)

type OCRService struct {
	extractor *extractors.ExtractionService
}

func NewOCRService() *OCRService {
	return &OCRService{
		extractor: extractors.NewExtractionService(),
	}
}

func (s *OCRService) ExtrairTexto(
	ctx context.Context,
	reader io.Reader,
) (*ports.OCRResult, error) {

	tmpFile :=
		os.TempDir() +
			"/" +
			uuid.NewString() +
			".pdf"

	file, err := os.Create(tmpFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	defer os.Remove(tmpFile)

	_, err = io.Copy(file, reader)
	if err != nil {
		return nil, err
	}

	return s.extractor.ExtrairTexto(
		ctx,
		tmpFile,
	)
}
