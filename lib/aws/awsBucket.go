package aws

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type SignedUrlResp struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type S3Bucket struct {
	Name   string
	Key    string
	Secret string
	Region string
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

func (bucket *S3Bucket) GetSignedUploadUrl() SignedUrlResp {
	// Prepare the S3 request so a signature can be generated
	sess := s3.New(session.New(&aws.Config{
		Region:      aws.String(bucket.Region),
		Credentials: credentials.NewStaticCredentials(bucket.Key, bucket.Secret, ""),
	}))

	key := generateStr(55) + ".pdf"
	r, _ := sess.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket.Name),
		Key:    aws.String(key),
	})

	// Create the pre-signed url with an expiry
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("Failed to generate a pre-signed url: ", err)
		return SignedUrlResp{}
	}

	// Display the pre-signed url
	return SignedUrlResp{
		Url: url,
		Key: key,
	}
}

func (bucket *S3Bucket) GetSignedReadUrl(key string) (string, error) {
	svc := s3.New(session.New(&aws.Config{
		Region:      aws.String(bucket.Region),
		Credentials: credentials.NewStaticCredentials(bucket.Key, bucket.Secret, ""),
	}))

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket.Name),
		Key:    aws.String(key),
	}

	req, _ := svc.GetObjectRequest(params)

	url, err := req.Presign(15 * time.Minute) // Set link expiration time
	if err != nil {
		log.Fatal("[AWS GET LINK]:", params, err)
	}

	return url, err
}

func (bucket *S3Bucket) GetPublicReadUrl(key string) string {
	url := "https://%s.s3.%s.amazonaws.com/%s"
	url = fmt.Sprintf(url, bucket.Name, bucket.Region, key)

	return url
}
