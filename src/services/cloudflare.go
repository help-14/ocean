package services

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	obconfig "github.com/help-14/ocean/config"
)

type CloudFlareService struct {
	name   string
	config obconfig.S3Config
	client *s3.Client
}

func (service *CloudFlareService) Name() string {
	return service.name
}

func (service *CloudFlareService) Setup(config obconfig.ServiceConfig) error {
	service.config = config.S3
	service.name = config.Name
	return nil
}

func (service *CloudFlareService) Connect() error {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(sv, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: service.config.Url,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(service.config.AccessKeyId, service.config.AccessKeySecret, "")),
	)
	if err != nil {
		return err
	}

	service.client = s3.NewFromConfig(cfg)
	return nil
}

func (service *CloudFlareService) Disconnect() error {
	service.client = nil
	return nil
}

func (service *CloudFlareService) Upload(localPath string, remotePath string) error {
	uploader := manager.NewUploader(service.client)
	uploadFile, err := os.Open(localPath)
	if err != nil {
		return err
	}

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(service.config.BucketName),
		Key:    aws.String(remotePath),
		Body:   uploadFile,
	})

	return err
}
