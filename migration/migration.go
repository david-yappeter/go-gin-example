package migration

import (
	"myapp/config"
	"myapp/entity"
)

func MigrateTable() {
	db := config.DB()

	db.AutoMigrate(&entity.User{})
}
