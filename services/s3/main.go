package s3

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"api-ngmi/lib/aws"
)

var AssetsBucket *aws.S3Bucket

var ReportsBucket *aws.S3Bucket

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	assetBucketName := os.Getenv("ASSET_BUCKET_NAME")
	if assetBucketName == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars asset")
	}

	assetBucketKey := os.Getenv("AWS_ACCESS_KEY_ID_ASSET")
	if assetBucketKey == "" {
		log.Fatal("no s3 key asset")
	}

	assetBucketSecret := os.Getenv("AWS_SECRET_ACCESS_KEY_ASSET")
	if assetBucketSecret == "" {
		log.Fatal("no s3 secret asset")
	}

	assetBucketRegion := os.Getenv("AWS_REGION_ASSET")
	if assetBucketRegion == "" {
		log.Fatal("no region asset")
	}

	AssetsBucket = &aws.S3Bucket{
		Name:   assetBucketName,
		Key:    assetBucketKey,
		Secret: assetBucketSecret,
		Region: assetBucketRegion,
	}

	reportBucketName := os.Getenv("REPORT_BUCKET_NAME")
	if reportBucketName == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars report")
	}

	reportBucketKey := os.Getenv("AWS_ACCESS_KEY_ID_REPORT")
	if reportBucketKey == "" {
		log.Fatal("no s3 key report")
	}

	reportBucketSecret := os.Getenv("AWS_SECRET_ACCESS_KEY_REPORT")
	if reportBucketSecret == "" {
		log.Fatal("no s3 secret report")
	}

	reportBucketRegion := os.Getenv("AWS_REGION_REPORT")
	if reportBucketRegion == "" {
		log.Fatal("no region report")
	}

	ReportsBucket = &aws.S3Bucket{
		Name:   reportBucketName,
		Key:    reportBucketKey,
		Secret: reportBucketSecret,
		Region: reportBucketRegion,
	}
}
