package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
