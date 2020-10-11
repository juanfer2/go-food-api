package repositories

import (
	"fmt"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func GetFoods() []models.Food {
	db := databases.Conn()
	var foods []models.Food
	db.Preload("TypeFood").Find(&foods)
	return foods
}

func GetFoodById(id interface{}) (food models.Food, err error) {
	db := databases.Conn()
	db.First(&food, id)

	if food.IsNull() {
		err = fmt.Errorf("food id: %s not found", id)
	}
	return food, err
}

func CreateFood(inputFood models.Food) models.Food {
	db := databases.Conn()
	food := inputFood
	db.Create(&food)

	return food
}

func UpdateFoodById(id interface{}, inputFood models.Food) (food models.Food, err error) {
	db := databases.Conn()
	db.First(&food, id)

	if food.IsNull() {
		err = fmt.Errorf("Food id: %s not found", id)
	}

	if !food.IsNull() {
		db.Model(&food).Where("id = ?", id).Updates(inputFood)
	}

	return food, err
}

func DeleteFoodById(id interface{}) (food models.Food, isDelete bool, err error) {
	db := databases.Conn()
	db.First(&food, id)

	if food.IsNull() {
		err = fmt.Errorf("food id: %s not found", id)
	}

	if !food.IsNull() {
		db.Delete(&food, id)
		isDelete = true
	}

	return food, isDelete, err
}
