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
	Produto     string       `json:"Produto"`
	Chamada     sql.NullTime `json:"senha chamada" gorm:"default=null"`
	Encerrada   sql.NullTime `json:"encerrada" gorm:"default=null"`
	Prioridade  bool         `json:"prioridade"`
}

type SenhasChamadas struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"criado"`
	UpdatedAt   time.Time    `json:"atualizado"`
	DeletedAt   time.Time    `json:"deletado"`
	Nome        string       `json:"nome"`
	SenhaGerada time.Time    `json:"senha gerada"`
	Produto     string       `json:"Produto"`
	Chamada     sql.NullTime `json:"senha chamada" gorm:"default=null"`
	Encerrada   sql.NullTime `json:"encerrada" gorm:"default=null"`
	Prioridade  bool         `json:"prioridade"`
}
