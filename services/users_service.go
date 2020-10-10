package services

import (
	"github.com/juanfer2/api-rest-go/constants"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateUserService(inputUser models.User) (models.User, error) {
	pass, _ := HashPassword(inputUser.Password)
	inputUser.Password = pass
	newUser := repositories.CreateUser(inputUser)
	token, _ := CreateToken(newUser.ID)
	newUser.Token = token
	newUser.SecretKey = constants.SecretKey
	return repositories.SaveUser(newUser)
}

func GetUsersService() []models.User {
	return repositories.GetUsers()
}
