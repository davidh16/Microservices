package initializers

import (
	"address-service/cmd/model"
	"log"
)

func init() {
	ConnectToDB()
}

func Migrate() {
	err := DB.AutoMigrate(&model.Address{})
	if err != nil {
		log.Fatalf("could not migrate to the database: %v", err)
	}
}
