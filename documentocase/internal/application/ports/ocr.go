package ports

import (
	"context"
	"io"
)

type OCRResult struct {
	Texto string
}

type Extractor interface {
	Extract(
		ctx context.Context,
		filePath string,
	) (*OCRResult, error)
}

type OCR interface {
	ExtrairTexto(
		ctx context.Context,
		arquivoID io.Reader,
	) (*OCRResult, error)
}
