package s3

import (
	"context"

	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AmazonStorageInterface interface {
	Upload(input *s3.PutObjectInput) (outPut *manager.UploadOutput, err error)
	Delete(input *s3.DeleteObjectInput) (outPut *s3.DeleteObjectOutput, err error)
}

type amazonStorage struct {
	buckets string
}

func NewAmazonStorage(bucket string) AmazonStorageInterface {
	return &amazonStorage{
		buckets: bucket,
	}
}

func (storage *amazonStorage) Upload(input *s3.PutObjectInput) (outPut *manager.UploadOutput, err error) {
	if len(storage.buckets) == 0 {

		//log.Error("S3 Bucket Name Can Not Be Empty")
		return nil, errors.New("S3 Bucket Name Can Not Be Empty")
	}

	input.Bucket = aws.String(storage.buckets)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {

		//log.Error(err)
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	return uploader.Upload(context.TODO(), input)
}

func (storage *amazonStorage) Delete(input *s3.DeleteObjectInput) (outPut *s3.DeleteObjectOutput, err error) {
	if len(storage.buckets) == 0 {
		return nil, errors.New("S3 Bucket Name Can Not Be Empty")
	}

	input.Bucket = aws.String(storage.buckets)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {

		//log.Error(err)
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	return client.DeleteObject(context.TODO(), input)
}
