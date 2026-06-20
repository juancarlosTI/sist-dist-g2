package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOStorage struct {
	client *minio.Client
	bucket string
}

func NewMinIOStorage(
	ctx context.Context,
	endpoint string,
	accessKey string,
	secretKey string,
	bucket string,
	useSSL bool,
) (*MinIOStorage, error) {

	if endpoint == "" {
		return nil, fmt.Errorf("minio endpoint não informado")
	}

	if bucket == "" {
		return nil, fmt.Errorf("minio bucket não informado")
	}

	client, err := minio.New(
		endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				accessKey,
				secretKey,
				"",
			),
			Secure: useSSL,
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"erro criando client minio: %w",
			err,
		)
	}

	exists, err := client.BucketExists(
		ctx,
		bucket,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"erro verificando bucket %s: %w",
			bucket,
			err,
		)
	}

	if !exists {

		err = client.MakeBucket(
			ctx,
			bucket,
			minio.MakeBucketOptions{},
		)

		if err != nil {

			// outra instância pode ter criado o bucket
			exists, bucketErr := client.BucketExists(
				ctx,
				bucket,
			)

			if bucketErr != nil || !exists {
				return nil, fmt.Errorf(
					"erro criando bucket %s: %w",
					bucket,
					err,
				)
			}
		}
	}

	return &MinIOStorage{
		client: client,
		bucket: bucket,
	}, nil
}
