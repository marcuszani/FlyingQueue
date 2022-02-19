package entities

import (
	"time"

	"gorm.io/gorm"
)

type AtendimentoQueue struct {
	gorm.Model
	Nome            string    `json:"nome"`
	DataAtendimento time.Time `json:"data atendimento"`
	Prioridade      bool      `json:"prioridade"`
}
