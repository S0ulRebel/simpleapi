package initializer

import (
	"simple-api/model"
)

func SyncDatabase() {
	PGDB.AutoMigrate(&model.User{})
	PGDB.AutoMigrate(&model.Post{})
}
