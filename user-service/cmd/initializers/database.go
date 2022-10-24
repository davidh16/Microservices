package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	dsn := os.Getenv("DSN")
	config := &gorm.Config{}
	counts := 0
	for {
		DB, err = gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
			counts++
		} else {
			log.Println("connected")
			return
		}

		if counts > 10 {
			log.Println(err)
			return
		}

		log.Println("waiting 2 sec before retry")
		time.Sleep(2 * time.Second)
		continue
	}
}
