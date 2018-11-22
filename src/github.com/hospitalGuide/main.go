package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hospitalGuide/controller"
)

func main(){
	r := gin.Default()

	lead := r.Group("/question") //导诊问题
	{
		lead.GET("/", controller.Question)
	}

	r.Run("0.0.0.0:9090") // 在 0.0.0.0:8080 上监听并服务
}
