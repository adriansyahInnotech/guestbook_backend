package db

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3Client *s3.S3

func InitS3Wasabi() {
	//
	// Membuat sesi AWS menggunakan kredensial Wasabi
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("WASABI_REGION")),   // Region Asia Pacific (Singapore)
		Endpoint:    aws.String(os.Getenv("WASABI_ENDPOINT")), // Endpoint Wasabi Asia
		Credentials: credentials.NewStaticCredentials(os.Getenv("WASABI_ACCESS_KEY"), os.Getenv("WASABI_SECRET_KEY"), os.Getenv("WASABI_TOKEN")),
	})
	if err != nil {
		log.Fatal("Failed to create session:", err)
	}

	S3Client = s3.New(sess)

}
