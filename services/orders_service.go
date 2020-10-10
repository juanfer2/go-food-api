package services

import (
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateOrderService(inputOrder models.Order) models.Order {
	return repositories.CreateOrder(inputOrder)
}

func GetOrdersService() []models.Order {
	return repositories.GetOrders()
}

func GetOrderByIdService(id interface{}) (order models.Order, err error) {
	order, err = repositories.GetOrderById(id)
	return order, err
}

func UpdateOrderByIdService(id interface{}, inputOrder models.Order) (order models.Order, err error) {
	order, err = repositories.UpdateOrderById(id, inputOrder)
	return order, err
}

func DeleteOrderByIdService(id interface{}) (order models.Order, isDelete bool, err error) {
	order, isDelete, err = repositories.DeleteOrderById(id)
	return order, isDelete, err
}
