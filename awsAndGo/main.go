package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	file, _ := ioutil.ReadFile("file1.txt")
	config := aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("KeyID", "SecretKey"),
	}

	s3Session := session.New(&config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("seeshasaimyfirstbucket"), // bucket's name
		Key:         aws.String("file1.txt"),              // files destination location
		Body:        bytes.NewReader(file),                // content of the file
		ContentType: aws.String("image/jpg"),              // content type
	}
	_, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		fmt.Println("error while uploading to the s3 bucket", err.Error())
	}
}
