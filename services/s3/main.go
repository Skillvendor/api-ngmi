package s3

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"api-ngmi/lib/aws"
)

var AssetsBucket *aws.S3Bucket

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	awsKey := os.Getenv("AWS_ACCESS_KEY_ID_CARPET")
	if s3Bucket == "" {
		log.Fatal("no s3 key")
	}

	awsSecret := os.Getenv("AWS_SECRET_ACCESS_KEY_CARPET")
	if s3Bucket == "" {
		log.Fatal("no s3 key")
	}

	awsRegion := os.Getenv("AWS_REGION_CARPET")
	if s3Bucket == "" {
		log.Fatal("no region")
	}

	AssetsBucket = &aws.S3Bucket{
		Name:   s3Bucket,
		Key:    awsKey,
		Secret: awsSecret,
		Region: awsRegion,
	}
}
