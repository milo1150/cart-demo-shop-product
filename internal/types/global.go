package types

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppState struct {
	DB    *gorm.DB
	Minio *minio.Client
	Log   *zap.Logger
}
