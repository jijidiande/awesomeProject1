package routers

import (
	"awesomeProject1/src/gin-blog/pkg/setting"
	v1 "awesomeProject1/src/gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}

//	router := routers.InitRouter()
//func InitRouter() *gin.Engine {
//	r := gin.New()
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//	gin.SetMode(setting.RunMode)
//
//	r.GET("/test", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "test",
//		})
//	})
//	return r
//}

//	router := gin.Default()
//Default包括	engine := New()	 engine.Use(Logger(), Recovery())
//	router.GET("/test", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "test",
//		})
//	})
