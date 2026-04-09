package s3app

import (
	"monorepo/shares/infra"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client() *S3Client {
	return &S3Client{
		client: infra.GetS3Client(
			"us-east-1",
			"saYp8zXw2rIbofww1L8V",
			"xAcAlyH4cZfk5IpPNMxlFGpHQpRW95yThINd8gmV",
			"http://localhost:9000",
		),
	}
}
