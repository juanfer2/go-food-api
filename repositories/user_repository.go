package repositories

import (
	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func GetUsers() []models.User {
	db := databases.Conn()
	var users []models.User
	db.Find(&users)
	return users
}

func CreateUser(inputUser models.User) models.User {
	db := databases.Conn()
	user := inputUser
	db.Create(&user)

	return user
}

func SaveUser(user models.User) (models.User, error) {
	db := databases.Conn()

	var err error

	if user.IsNull() {
		save_user := db.Create(&user)
		if save_user.Error != nil {
			err = save_user.Error
		}
	} else {
		save_user := db.Save(&user)
		if save_user.Error != nil {
			err = save_user.Error
		}
	}

	return user, err
}
