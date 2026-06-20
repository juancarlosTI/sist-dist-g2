package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (s *MinIOStorage) Delete(
	ctx context.Context,
	fileID string,
) error {

	return s.client.RemoveObject(
		ctx,
		s.bucket,
		fileID,
		minio.RemoveObjectOptions{},
	)
}
