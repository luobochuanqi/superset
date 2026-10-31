package model

type Model struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;comment:主键，自增"`
	CreatedAt int64  `gorm:"not null;autoCreateTime;comment:创建时间"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime;comment:更新时间"`
	DeletedAt int64  `gorm:"null;comment:软删除：非 NULL 即视为已删除（GORM 约定）"`
}
