package services

import (
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/adapter"
	"gorm.io/gorm"
)

type Service struct {
	DB           *gorm.DB
	RedisAdapter *adapter.RedisAdapter
}

func NewService(db *gorm.DB, ra *adapter.RedisAdapter) *Service {
	return &Service{
		DB:           db,
		RedisAdapter: ra,
	}
}
