package model

type Article struct {
	*Model
	Title         string `gorm:"type:varchar(256);not null;comment:标题"`
	Content       string `gorm:"type:text;not null;comment:正文"`
	CoverImageURL string `gorm:"type:varchar(512);not null;default:'';comment:封面图"`
}

func (Article) TableName() string {
	return "blog_article"
}
