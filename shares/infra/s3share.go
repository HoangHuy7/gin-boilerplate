package infra

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Client(
	region string,
	access_key_id string,
	secret_access_key string,
	endpoint string,
) *s3.Client {
	if access_key_id == "" || secret_access_key == "" || region == "" || endpoint == "" {
		log.Fatal("missing the env: RUSTFS_ACCESS_KEY_ID / RUSTFS_SECRET_ACCESS_KEY / RUSTFS_REGION / RUSTFS_ENDPOINT_URL")
	}
	cfg := aws.Config{
		Region: region,
		EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: endpoint,
			}, nil
		}),
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(access_key_id, secret_access_key, "")),
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	_, err := client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String("go-sdk-rustfs"),
	})
	if err != nil {
		log.Fatalf("create bucket failed: %v", err)
	}

	resp, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String("mekyrax"),
		Key:    aws.String("phan1.png"),
	})
	if err != nil {
		log.Fatalf("download object fail: %v", err)
	}
	defer resp.Body.Close()

	// read object content
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read object content fail: %v", err)
	}
	fmt.Println("content is :", string(data))

	return client
}
