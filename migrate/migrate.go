package migrate

import (
	"github.com/Asliddin3/tz/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Person{},
	)
	if err != nil {
		return err
	}

	return nil
}
