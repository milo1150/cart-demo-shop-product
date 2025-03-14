package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

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

func CreateBucket(client *minio.Client, ctx context.Context, bucketName string) {
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("MinIO: already create %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("MinIO: successfully created %s\n", bucketName)
	}
}

func UploadFile(client *minio.Client, ctx context.Context, bucketName, objectName, filePath, contentType string, log *zap.Logger) {
	info, err := client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error(fmt.Sprintf("Failed to upload file %s to %s", objectName, bucketName), zap.Error(err))
	}
	log.Info(fmt.Sprintf("Successfully uploaded %s of size %d", objectName, info.Size))
}
