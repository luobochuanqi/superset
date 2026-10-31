package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// 增删改查
func (Tag) Create(ctx *gin.Context) {}
func (Tag) Update(ctx *gin.Context) {}
func (Tag) Delete(ctx *gin.Context) {}
func (Tag) Get(ctx *gin.Context)    {}
func (Tag) List(ctx *gin.Context)   {}
