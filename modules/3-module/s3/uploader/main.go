package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/*
   https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/getting-started.html
*/

var (
	s3Client *s3.Client
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithSharedConfigProfile("rcmelo-vscode"),
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.NewFromConfig(cfg)
	s3Bucket = "goexpert-s3-example-bucket"
}

func main() {
	dir, err := os.Open("../tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	// Retry failed uploads
	go func() {
		for filename := range errorFileUpload {
			uploadControl <- struct{}{}
			wg.Add(1)
			go uploadFile(filename, uploadControl, errorFileUpload)
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("../tmp/%s", filename)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		// empty the channel
		<-uploadControl
		errorFileUpload <- filename
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", filename, err)
		<-uploadControl
		errorFileUpload <- filename
		return
	}
	fmt.Printf("Successfully uploaded %s to %s\n", filename, s3Bucket)
	<-uploadControl
}
