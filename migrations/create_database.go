package migrations

import (
	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func main() {
	db := databases.Conn()
	db.AutoMigrate(&models.Task{})
}
