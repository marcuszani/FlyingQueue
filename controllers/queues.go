package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FlyingQueue/entities"
	"github.com/FlyingQueue/models"
	"github.com/gin-gonic/gin"
)

func init() {

}

func NovoAtendimento(c *gin.Context) {

	novoRegistro := entities.AtendimentoQueue{
		Nome:            c.Param("nome"),
		DataAtendimento: time.Now(),
		Prioridade:      false,
	}

	err := models.NovoAtendimento(novoRegistro)

	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Registro inserido com sucesso",
		})
	}

}

func BuscarTodosAtendimentos(c *gin.Context) {

	c.JSON(200, models.TodosAtendimentos())

}

func BuscarAtendimentoPorID(c *gin.Context) {

	registro := models.BuscarAtendimentoPorID(c.Params.ByName("id"))

	if registro.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "Registro não encontrado",
		})
		return
	} else {
		c.JSON(http.StatusOK, registro)
	}

}
