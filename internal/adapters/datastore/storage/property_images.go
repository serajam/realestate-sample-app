/* Copyright (C) Fedir Petryk */

package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
)

type Images struct {
	client   *minio.Client
	bucket   string
	location string
}

func NewImagesStorage(ctx context.Context, client *minio.Client, bucketName, location string) (*Images, error) {
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists != nil || !exists {
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	policy := fmt.Sprintf(
		`{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},
"Resource": ["arn:aws:s3:::%s/*"],"Sid": ""}]}`, bucketName,
	)

	err = client.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		return nil, err
	}

	return &Images{client: client, bucket: bucketName, location: location}, nil
}

func (s *Images) Upload(ctx context.Context, img *properties.Image) error {
	contentType := "image/jpeg"

	for _, content := range img.Contents {

		// Upload the zip file with FPutObject
		_, err := s.client.PutObject(
			ctx, s.bucket, properties.ImageObjectName(img.ID.String(), content.ImageType), bytes.NewReader(content.Content.Bytes()),
			int64(content.Content.Len()),
			minio.PutObjectOptions{ContentType: contentType},
		)
		if err != nil {
			return err
		}

	}

	return nil
}

func (s *Images) Get(ctx context.Context, name string) (*minio.Object, error) {
	info, err := s.client.GetObject(ctx, s.bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (s *Images) Del(ctx context.Context, name string) error {
	err := s.client.RemoveObject(ctx, s.bucket, name, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
