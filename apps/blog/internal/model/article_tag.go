package model

type ArticleTag struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;comment:主键，自增"`
	ArticleID uint64 `gorm:"not null;comment:属于哪篇文章"`
	TagID     uint64 `gorm:"not null;comment:属于哪个标签"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
