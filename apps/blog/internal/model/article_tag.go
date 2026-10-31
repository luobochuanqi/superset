package model

type ArticleTag struct {
	Model
	ArticleID uint64 `gorm:"not null;comment:属于哪篇文章"`
	TagID     uint64 `gorm:"not null;comment:属于哪个标签"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
