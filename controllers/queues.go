package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	NewQueue()
}

type Queue struct {
	Name string `json:"name"`
}

var List []Queue

func NewQueue() {

	tmpList := Queue{
		Name: "Marcus",
	}

	tmpList1 := Queue{
		Name: "Carlos",
	}

	List = append(List, tmpList, tmpList1)

	fmt.Println(List)

}

func ReadQueue(c *gin.Context) {

	c.JSON(200, List)

}
