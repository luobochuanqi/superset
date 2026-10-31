package main

import (
	router "github.com/luobochuanqi/superset/apps/blog/internal/routers"
	"net/http"
)

func main() {
	// r := gin.Default()
	router := router.NewRouter()
	s := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	s.ListenAndServe()

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "pong"})
	// })
	// r.Run()
}
