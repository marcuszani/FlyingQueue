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
		Nome:        c.Param("nome"),
		SenhaGerada: time.Now(),
		Prioridade:  false,
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

func DeletarAtendimentoPorID(c *gin.Context) {
	err := models.DeletarAtendimentoPorID(c.Params.ByName("id"))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Falha ao deletar registro",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Registro deletado com sucesso",
		})
	}

}

func AtualizarAtendimentoPorID(c *gin.Context) {

	novoNome := "teste"

	err := models.AtualizarAtendimentoPorID(c.Param("id"), novoNome)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Problema ao atualizar registro",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Registro alterado com sucesso",
		})
	}

}

func ChamarSenha(c *gin.Context) {

	senha, contagem := models.ChamarSenha()

	if contagem > 0 {
		c.JSON(http.StatusOK, senha)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Nenhuma senha disponível",
		})
	}

}
