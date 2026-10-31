package model

type Tag struct {
	Model
	Name string `gorm:"type:varchar(64);not null;uniqueIndex:uk_name;comment:标签名"`
}

func (Tag) TableName() string {
	return "blog_tag"
}
