package extractors

// import (
// 	"context"

// 	"github.com/otiai10/gosseract/v2"

// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
// )

// type TesseractExtractor struct{}

// func NewTesseractExtractor() *TesseractExtractor {
// 	return &TesseractExtractor{}
// }

// func (e *TesseractExtractor) Extract(
// 	ctx context.Context,
// 	filePath string,
// ) (*ports.OCRResult, error) {

// 	select {
// 	case <-ctx.Done():
// 		return nil, ctx.Err()
// 	default:
// 	}

// 	client := gosseract.NewClient()
// 	defer client.Close()

// 	err := client.SetImage(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	texto, err := client.Text()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ports.OCRResult{
// 		Texto: texto,
// 	}, nil
// }
