package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/services"
)

func parserBodyOrder(c *fiber.Ctx) models.Order {
	newOrder := new(models.Order)
	c.BodyParser(newOrder)
	return *newOrder
}

func CreateOrder(c *fiber.Ctx) error {
	var inputOrder models.Order
	inputOrder = parserBodyOrder(c)

	data := services.CreateOrderService(inputOrder)
	return c.JSON(data)
}

func GetOrders(c *fiber.Ctx) error {
	orders := services.GetOrdersService()
	return c.JSON(orders)

}

func GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := services.GetOrderByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(order)
}

func UpdateOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	var inputOrder models.Order
	inputOrder = parserBodyOrder(c)
	order, err := services.UpdateOrderByIdService(id, inputOrder)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(order)
}

func DeleteOrderById(c *fiber.Ctx) error {
	id := c.Params("id")
	msg := ""
	_, isDelete, err := services.DeleteOrderByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	if isDelete {
		msg = fmt.Sprintf("Order id %s has been deleted", id)
	} else {
		msg = fmt.Sprintf("Order id %s has not been deleted", id)
	}

	return c.JSON(msg)
}
