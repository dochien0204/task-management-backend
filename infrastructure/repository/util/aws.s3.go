package util

import (
	"context"

	otherConfig "source-base-go/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func GeneratePresignURLS3(keyName string) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(client)

	presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(otherConfig.S3_BUCKET_NAME),
		Key:    aws.String(keyName),
	})

	if err != nil {
		return "", err
	}

	return presignResult.URL, nil
}

func GeneratePresignUploadS3(keyName string) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(client)

	presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(otherConfig.S3_BUCKET_NAME),
		Key:    aws.String(keyName)},
	)

	if err != nil {
		return "", err
	}

	return presignResult.URL, nil
}