package migrations

import (
	"github.com/FlyingQueue/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		models.AtendimentoQueue{},
	)

}
