package storage

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

func (s *MinIOStorage) Download(
	ctx context.Context,
	fileID string,
) (io.ReadCloser, error) {

	obj, err :=
		s.client.GetObject(
			ctx,
			s.bucket,
			fileID,
			minio.GetObjectOptions{},
		)

	if err != nil {
		return nil, err
	}

	_, err = obj.Stat()
	if err != nil {
		return nil, err
	}

	return obj, nil
}
