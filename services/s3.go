package services

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getConfig() s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return s3.Client{}
	}

	client := s3.NewFromConfig(cfg)
	return *client
}

func UploadToBucket(bucketName string, filename string) {
	client := getConfig()
	objectKey := filename

	filePath := "tmp_files/" + filename
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   file,
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("File uploaded successfully!")
}

func Download(bucketName string, filename string) string {
	client := getConfig()
	result, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &filename,
	})

	if err != nil {
		fmt.Println("Error downloading file:", err)
	}

	defer result.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	_, err = file.Write(body)

	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	return file.Name()
}

func GetSignedUrl(bucketName string, filename string) string {
	client := getConfig()

	presignClient := s3.NewPresignClient(&client)

	url, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &filename,
	})

	if err != nil {
		fmt.Println("Error presigning request:", err)
		return ""
	}

	return url.URL
}
