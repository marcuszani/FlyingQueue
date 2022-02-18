package routes

import (
	"github.com/FlyingQueue/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/status", controllers.StatusServer)
	r.GET("/list", controllers.ReadQueue)

	r.Run(":8080")
}
