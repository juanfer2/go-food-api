package databases

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  "user=postgres password=1234 dbname=go-api port=5432 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
			PreferSimpleProtocol: true,                                                                  // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	log.Println("Database connected")

	return db
}
