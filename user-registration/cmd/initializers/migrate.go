package initializers

import (
	"log"
	"user-registration/cmd/model"
)

func init() {
	ConnectToDB()
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("could not migrate to the database: %v", err)
	}
}
