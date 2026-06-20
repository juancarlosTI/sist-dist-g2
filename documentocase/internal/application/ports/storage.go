package ports

import (
	"context"
	"io"
)

type UploadedFile struct {
	ID       string
	Filename string
	Size     int64
}

type Storage interface {
	Upload(
		ctx context.Context,
		reader io.Reader,
		filename string,
		contentType string,
	) (*UploadedFile, error)

	Download(
		ctx context.Context,
		fileID string,
	) (io.ReadCloser, error)

	Delete(
		ctx context.Context,
		fileID string,
	) error
}
