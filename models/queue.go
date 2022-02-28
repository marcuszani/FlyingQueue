package models

import (
	"fmt"
	"time"

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

func DeletarAtendimentoPorID(id string) error {
	db := database.GetDatabase()

	atendimento := entities.AtendimentoQueue{}

	err := db.Where("id = ?", id).Delete(atendimento).Error

	return err
}

func AtualizarAtendimentoPorID(id, nome string) error {
	db := database.GetDatabase()

	atendimento := entities.AtendimentoQueue{}

	err := db.Model(&atendimento).Where("id = ?", id).Update("nome", nome).Error

	return err
}

func ChamarSenha() (*entities.AtendimentoQueue, int64) {

	senhaChamada := entities.AtendimentoQueue{}

	db := database.GetDatabase()

	contagem := db.Where("chamada IS ? AND encerrada is ? ", nil, nil).First(&senhaChamada).RowsAffected

	if contagem > 0 {
		db.Model(&senhaChamada).Where("id = ?", senhaChamada.ID).Update("chamada", time.Now())

	}

	return &senhaChamada, contagem

}
func EncerrarAtendimento(senha string) error {
	db := database.GetDatabase()

	senhaChamada := entities.AtendimentoQueue{}

	err := db.Model(&senhaChamada).Where("id = ?", senha).Update("encerrada", time.Now()).Error

	if err != nil {
		fmt.Println(err)
	} else {
		backupSenha := entities.SenhasChamadas{
			ID:          senhaChamada.ID,
			CreatedAt:   senhaChamada.CreatedAt,
			UpdatedAt:   senhaChamada.UpdatedAt,
			DeletedAt:   senhaChamada.DeletedAt.Time,
			Nome:        senhaChamada.Nome,
			SenhaGerada: senhaChamada.SenhaGerada,
			Chamada:     senhaChamada.Chamada,
			Encerrada:   senhaChamada.Encerrada,
			Prioridade:  senhaChamada.Prioridade,
		}
		fmt.Println(backupSenha)

		err = db.Create(&backupSenha).Error
	}

	return err
}
