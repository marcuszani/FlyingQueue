package routes

import (
	"github.com/FlyingQueue/config"
	"github.com/FlyingQueue/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.SetTrustedProxies([]string{config.Cfg.Server["Url"]})

	r.GET("/status", controllers.StatusServer)
	r.GET("/list", controllers.ReadQueue)
	r.GET("/new", controllers.NewQueue)

	r.Run(":" + config.Cfg.Server["Port"])
}
