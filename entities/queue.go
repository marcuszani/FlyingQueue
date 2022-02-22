package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type AtendimentoQueue struct {
	gorm.Model
	Nome        string       `json:"nome"`
	SenhaGerada time.Time    `json:"senha gerada"`
	Chamada     sql.NullTime `json:"senha chamada" gorm:"default=null"`
	Encerrada   sql.NullTime `json:"encerrada" gorm:"default=null"`
	Prioridade  bool         `json:"prioridade"`
}
