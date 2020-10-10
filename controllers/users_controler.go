package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/services"
	"github.com/juanfer2/api-rest-go/validations"
)

type ErrorMessage string

func parserBodyUser(c *fiber.Ctx) models.User {
	newUser := new(models.User)
	c.BodyParser(newUser)
	return *newUser
}

// Create Task
func CreateUser(c *fiber.Ctx) error {
	var inputUser models.User
	inputUser = parserBodyUser(c)
	err := validations.ValidateCreateUser(inputUser)

	if err != nil {
		return c.Status(422).JSON(err)
	}

	data, err_save := services.CreateUserService(inputUser)

	if err_save != nil {
		msg := fmt.Sprint(err)
		return c.Status(422).JSON(msg)
	}
	return c.JSON(data)
}

func GetUsers(c *fiber.Ctx) error {
	users := services.GetUsersService()
	return c.JSON(users)
}
