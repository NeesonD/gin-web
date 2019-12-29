package routers

import (
	"gin-web/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message": "hello",
		})
	})
	return r
}