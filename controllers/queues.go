package controllers

import (
	"fmt"
	"time"

	"github.com/FlyingQueue/database"
	"github.com/FlyingQueue/models"
	"github.com/gin-gonic/gin"
)

func init() {

}

type Queue struct {
	Name string `json:"name"`
}

var List []Queue

func NewQueue(c *gin.Context) {

	novoRegistro := models.AtendimentoQueue{
		DataAtendimento: time.Now(),
		Prioridade:      false,
	}

	db := database.GetDatabase()

	db.Create(&novoRegistro)
	c.JSON(200, gin.H{
		"status": "Registro inserido com sucesso",
	})

}

func ReadQueue(c *gin.Context) {

	todosRegistros := []models.AtendimentoQueue{}

	db := database.GetDatabase()

	err := db.Find(&todosRegistros).Error

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, todosRegistros)

}
