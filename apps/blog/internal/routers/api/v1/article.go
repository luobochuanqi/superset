package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// 增删改查
func (Article) Create(ctx *gin.Context) {}
func (Article) Update(ctx *gin.Context) {}
func (Article) Delete(ctx *gin.Context) {}
func (Article) Get(ctx *gin.Context)    {}
func (Article) List(ctx *gin.Context)   {}
