package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/luobochuanqi/superset/apps/blog/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	{
		/**
		新增标签		POST	/tags
		删除指定标签	DELETE	/tags/:id
		更新指定标签	PUT		/tags/:id
		获取标签列表	GET		/tags
		*/
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.GET("/tags", tag.List)

		/**
		新增文章		POST	/articles
		删除指定文章	DELETE	/articles/:id
		更新指定文章	PUT		/articles/:id
		获取指定文章	GET		/articles/:id
		获取文章列表	GET		/articles
		*/
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
