package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	s3session *s3.S3
)

const (
	BUCKETNAME = "20201108-bucket2"
	REGION     = "ap-northeast-1"
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func listBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	return resp
}

func createBuckets() (resp *s3.CreateBucketOutput) {
	resp, err := s3session.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(BUCKETNAME),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(REGION),
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println("Bucket Name already in use.")
				panic(err)
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println("BUckey exists and is owned by you.")
			default:
				panic(err)
			}
		}
	}
	return resp
}

func uploadHandler(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	fileName := file.Filename

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstfile, err := os.Create("./tmp/" + fileName)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	if _, err = io.Copy(dstfile, src); err != nil {
		return err
	}
	uploadObject("tmp/" + fileName)
	return nil
}
func uploadObject(filename string) (resp *s3.PutObjectOutput) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploading:", filename)
	resp, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(BUCKETNAME),
		Key:    aws.String(strings.Split(filename, "/")[1]),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		panic(err)
	}
	return resp
}

func main() {
	// uploadObject("img/rails.png")

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/upload", uploadHandler)

	e.Start(":9999")
}
