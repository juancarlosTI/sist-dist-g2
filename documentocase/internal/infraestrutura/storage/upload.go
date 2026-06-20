package storage

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

func (s *MinIOStorage) Upload(
	ctx context.Context,
	reader io.Reader,
	filename string,
	contentType string,
) (*ports.UploadedFile, error) {

	id := uuid.NewString()

	info, err :=
		s.client.PutObject(
			ctx,
			s.bucket,
			id,
			reader,
			-1,
			minio.PutObjectOptions{
				ContentType: contentType,
			},
		)

	if err != nil {
		return nil, err
	}

	return &ports.UploadedFile{
		ID:       id,
		Filename: filename,
		Size:     info.Size,
	}, nil
}
