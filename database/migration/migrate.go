package migration

import (
	"fmt"
	"log"

	"github.com/mawitra/test/database"
	"github.com/mawitra/test/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Author{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
