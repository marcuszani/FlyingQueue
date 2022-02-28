package entities

import (
	"time"

	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome      string    `json:"Nome"`
	Prefixo   string    `json:"Prefixo"`
	Categoria Categoria `json:"Categoria"`
}

type Categoria struct {
	gorm.Model
	Nome string `json:"Nome"`
}

type ProdutosVendidos struct {
	Data     time.Time `json:"Data"`
	Produto  Produto   `json:"Produto"`
	Contagem uint      `json:"Contagem"`
}
