package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/services"
)

func parserBodyTypeFood(c *fiber.Ctx) models.TypeFood {
	newTypeFood := new(models.TypeFood)
	c.BodyParser(newTypeFood)
	return *newTypeFood
}

func CreateTypeFood(c *fiber.Ctx) error {
	var inputTypeFood models.TypeFood
	inputTypeFood = parserBodyTypeFood(c)
	data := services.CreateTypeFoodService(inputTypeFood)
	return c.JSON(data)
}

func GetTypeFoods(c *fiber.Ctx) error {
	typeFoods := services.GetTypeFoodsService()
	return c.JSON(typeFoods)

}

func GetTypeFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	typeFood, err := services.GetTypeFoodByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(typeFood)
}

func UpdateTypeFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	var inputTypeFood models.TypeFood
	inputTypeFood = parserBodyTypeFood(c)
	typeFood, err := services.UpdateTypeFoodByIdService(id, inputTypeFood)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(typeFood)
}

func DeleteTypeFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	msg := ""
	_, isDelete, err := services.DeleteTypeFoodByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	if isDelete {
		msg = fmt.Sprintf("TypeFood id %s has been deleted", id)
	} else {
		msg = fmt.Sprintf("TypeFood id %s has not been deleted", id)
	}

	return c.JSON(msg)
}
