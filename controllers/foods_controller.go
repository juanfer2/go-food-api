package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"

	"github.com/juanfer2/api-rest-go/services"
)

type FoodResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func parserBodyFood(c *fiber.Ctx) models.Food {
	newFood := new(models.Food)
	c.BodyParser(newFood)
	return *newFood
}

func CreateFood(c *fiber.Ctx) error {
	var inputFood models.Food
	inputFood = parserBodyFood(c)
	data := services.CreateFoodService(inputFood)
	return c.JSON(data)
}

func GetFoods(c *fiber.Ctx) error {
	foods := services.GetFoodsService()
	return c.JSON(fiber.Map{"status": "success", "message": "All foods", "data": foods})
}

func GetFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	food, err := services.GetFoodByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(food)
}

func UpdateFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	var inputFood models.Food
	inputFood = parserBodyFood(c)
	food, err := services.UpdateFoodByIdService(id, inputFood)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(food)
}

func DeleteFoodById(c *fiber.Ctx) error {
	id := c.Params("id")
	msg := ""
	_, isDelete, err := services.DeleteFoodByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	if isDelete {
		msg = fmt.Sprintf("Food id %s has been deleted", id)
	} else {
		msg = fmt.Sprintf("Food id %s has not been deleted", id)
	}

	return c.JSON(msg)
}

/*
{
  "printWidth": 80,
  "tabWidth": 2,
  "useTabs": false,
  "semi": false,
  "singleQuote": true,
  "traillingComma": "all",
  "bracketSpacing": false,
  "proseWrap": "always",
  "jsxBracketSameLine": false
}
*/
