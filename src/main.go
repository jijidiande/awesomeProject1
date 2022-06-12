package main

import (
	"awesomeProject1/src/gin-blog/routers"
	"fmt"
	"net/http"

	"awesomeProject1/src/gin-blog/pkg/setting"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	//
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
