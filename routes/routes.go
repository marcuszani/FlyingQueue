package routes

import (
	"github.com/FlyingQueue/config"
	"github.com/FlyingQueue/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.SetTrustedProxies([]string{config.Cfg.Server["Url"]})

	r.Static("/static", "./assets")

	r.GET("/status", controllers.StatusServer)

	gpAtendimento := r.Group("/atendimento")
	{
		gpAtendimento.GET("/listar", controllers.BuscarTodosAtendimentos)
		gpAtendimento.GET("/novo", controllers.NovoAtendimento) //alterar para POST
		gpAtendimento.GET("/:id", controllers.BuscarAtendimentoPorID)
		gpAtendimento.DELETE("/:id", controllers.DeletarAtendimentoPorID)
		gpAtendimento.GET("/chamar", controllers.ChamarSenha)
		gpAtendimento.GET("/:id/encerrar", controllers.EncerrarAtendimento)

	}

	r.Run(":" + config.Cfg.Server["Port"])
}
