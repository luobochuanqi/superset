package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;comment:主键，自增"`
	CreatedAt time.Time      `gorm:"not null;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"comment:软删除：非 NULL 即视为已删除（GORM 约定）"`
}
