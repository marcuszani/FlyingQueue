package models

import (
	"time"

	"gorm.io/gorm"
)

type AtendimentoQueue struct {
	gorm.Model
	DataAtendimento time.Time `json:"data atendimento"`
	Prioridade      bool      `json: "prioridade"`
}

// func TodosAtendimentos() {
// 	db := database.GetDatabase()

// 	atendimentos := AtendimentoQueue{}

// 	todosAtendimentos := db.Find(&atendimentos)

// 	fmt.Println(todosAtendimentos)

// }
