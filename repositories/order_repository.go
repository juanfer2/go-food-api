package repositories

import (
	"fmt"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func GetOrders() []models.Order {
	db := databases.Conn()
	var orders []models.Order
	db.Find(&orders)
	return orders
}

func GetOrderById(id interface{}) (order models.Order, err error) {
	db := databases.Conn()
	db.First(&order, id)

	if order.IsNull() {
		err = fmt.Errorf("Order id: %s not found", id)
	}
	return order, err
}

func CreateOrder(inputOrder models.Order) models.Order {
	db := databases.Conn()
	order := inputOrder
	db.Create(&order)

	return order
}

func UpdateOrderById(id interface{}, inputOrder models.Order) (order models.Order, err error) {
	db := databases.Conn()
	db.First(&order, id)

	if order.IsNull() {
		err = fmt.Errorf("Order id: %s not found", id)
	}

	if !order.IsNull() {
		db.Model(&order).Where("id = ?", id).Updates(inputOrder)
	}

	return order, err
}

func DeleteOrderById(id interface{}) (order models.Order, isDelete bool,
	err error) {
	db := databases.Conn()
	db.First(&order, id)

	if order.IsNull() {
		err = fmt.Errorf("Order id: %s not found", id)
	}

	if !order.IsNull() {
		db.Delete(&order, id)
		isDelete = true
	}

	return order, isDelete, err
}
