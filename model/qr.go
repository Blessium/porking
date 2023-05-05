package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type QR struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	Image     string
}

func (q *QR) BeforeCreate(scope *gorm.DB) error {
	q.ID = uuid.NewV4().String()
	return nil
}
