package s3

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

type SignedUrlResp struct {
	Url string
	Key string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateStr(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func GetSignedUploadUrl() SignedUrlResp {
	// Load the bucket name
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

	// Prepare the S3 request so a signature can be generated
	sess := s3.New(session.New(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION_CARPET")),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, ""),
	}))

	key := generateStr(55) + ".png"
	r, _ := sess.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(key),
	})

	// Create the pre-signed url with an expiry
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("Failed to generate a pre-signed url: ", err)
		return SignedUrlResp{}
	}

	log.Println("URL", url)

	// Display the pre-signed url
	return SignedUrlResp{
		Url: url,
		Key: key,
	}
}

func GetSignedReadUrl(key string) (string, error) {
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

	svc := s3.New(session.New(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION_CARPET")),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, ""),
	}))

	params := &s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(key),
	}

	req, _ := svc.GetObjectRequest(params)

	url, err := req.Presign(15 * time.Minute) // Set link expiration time
	if err != nil {
		log.Fatal("[AWS GET LINK]:", params, err)
	}

	return url, err
}

func GetPublicReadUrl(key string) string {
	region := os.Getenv("AWS_REGION_CARPET")
	if region == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	url := "https://%s.s3.%s.amazonaws.com/%s"
	url = fmt.Sprintf(url, s3Bucket, region, key)

	return url
}
