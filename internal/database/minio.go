package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type MinIO struct {
	Client  *minio.Client
	Context context.Context
	ApiURL  string
	Log     *zap.Logger
}

func ConnectMinioDatabase() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	return minioClient
}

func (m *MinIO) CreateBucket(bucketName string) {
	err := m.Client.MakeBucket(m.Context, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.Client.BucketExists(m.Context, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("MinIO: already create %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("MinIO: successfully created %s\n", bucketName)
	}
}

// True if file already exists
func (m *MinIO) FileExists(bucketName, objectName string) bool {
	_, err := m.Client.StatObject(m.Context, bucketName, objectName, minio.GetObjectOptions{})
	return err == nil
}

func (m *MinIO) UploadFile(bucketName, objectName, filePath, contentType string, log *zap.Logger) {
	info, err := m.Client.FPutObject(m.Context, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error(fmt.Sprintf("Failed to upload file %s to %s", objectName, bucketName), zap.Error(err))
	}

	// Log
	fileURL := fmt.Sprintf("%s/%s/%s", m.Client.EndpointURL(), bucketName, objectName)
	log.Info("File accessible at:", zap.String("URL", fileURL))
	log.Info(fmt.Sprintf("Successfully uploaded %s of size %d", objectName, info.Size))
}

func (m *MinIO) GetPublicURL(bucketName, objectName string) string {
	return fmt.Sprintf("%s/%s/%s", m.ApiURL, bucketName, objectName)
}

func (m *MinIO) GetPublicURLWithExpireDate(bucketName, objectName string, expires time.Duration) string {
	presignedUrl, err := m.Client.PresignedGetObject(m.Context, bucketName, objectName, expires, nil)
	if err != nil {
		m.Log.Error("Error GeneratePublicURLWithExpireDate", zap.Error(err))
	}

	return fmt.Sprintf("%s%s?%s", m.ApiURL, presignedUrl.Path, presignedUrl.RawQuery)
}

func CreatePublicBucket(client *minio.Client, bucketName string) {
	ctx := context.Background()
	location := "us-east-1"

	// Create the bucket if it doesn't exist
	exists, errBucketExists := client.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		log.Fatalln(errBucketExists)
	}
	if !exists {
		err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Set public read-only policy on the bucket
	policy := `{
	    "Version": "2012-10-17",
	    "Statement": [
	        {
	            "Effect": "Allow",
	            "Principal": "*",
	            "Action": ["s3:GetObject"],
	            "Resource": ["arn:aws:s3:::` + bucketName + `/*"]
	        }
	    ]
	}`

	err := client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		log.Fatalln("Error setting bucket policy:", err)
	}

	log.Println("Bucket is now public!")
}
