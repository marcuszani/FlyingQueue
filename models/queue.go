package models

import (
	"fmt"

	"github.com/FlyingQueue/database"
	"github.com/FlyingQueue/entities"
)

// type AtendimentoQueue struct {
// 	gorm.Model
// 	DataAtendimento time.Time `json:"data atendimento"`
// 	Prioridade      bool      `json:"prioridade"`
// }

func TodosAtendimentos() []entities.AtendimentoQueue {
	db := database.GetDatabase()

	atendimentos := []entities.AtendimentoQueue{}

	err := db.Find(&atendimentos).Error

	if err != nil {
		fmt.Println(err)
	}

	return atendimentos
}

func NovoAtendimento(novoRegistro entities.AtendimentoQueue) error {

	db := database.GetDatabase()

	err := db.Create(&novoRegistro).Error

	return err

}

func BuscarAtendimentoPorID(id string) entities.AtendimentoQueue {
	db := database.GetDatabase()

	atendimento := entities.AtendimentoQueue{}

	err := db.Where("id = ?", id).First(&atendimento)

	if err != nil {
		fmt.Println(err)
	}

	return atendimento
}
