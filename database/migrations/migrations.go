package migrations

import (
	"github.com/FlyingQueue/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		entities.AtendimentoQueue{},
		entities.SenhasChamadas{},
		entities.Categoria{},
		entities.Produto{},
	)

}
