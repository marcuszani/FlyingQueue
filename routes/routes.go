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

	gpAtendimento := r.Group("/atendimento")
	{
		gpAtendimento.GET("/listar", controllers.BuscarTodosAtendimentos)
		gpAtendimento.POST("/novo", controllers.NovoAtendimento)
		gpAtendimento.GET("/:id", controllers.BuscarAtendimentoPorID)

	}

	r.Run(":" + config.Cfg.Server["Port"])
}
