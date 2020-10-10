package repositories

import (
	"fmt"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func GetTypeFoods() []models.TypeFood {
	db := databases.Conn()
	var typeFoods []models.TypeFood
	db.Find(&typeFoods)
	return typeFoods
}

func GetTypeFoodById(id interface{}) (typeFood models.TypeFood, err error) {
	db := databases.Conn()
	db.First(&typeFood, id)

	if typeFood.IsNull() {
		err = fmt.Errorf("TypeFood id: %s not found", id)
	}
	return typeFood, err
}

func CreateTypeFood(inputTypeFood models.TypeFood) models.TypeFood {
	db := databases.Conn()
	typeFood := inputTypeFood
	db.Create(&typeFood)

	return typeFood
}

func UpdateTypeFoodById(id interface{}, inputTypeFood models.TypeFood) (typeFood models.TypeFood, err error) {
	db := databases.Conn()
	db.First(&typeFood, id)

	if typeFood.IsNull() {
		err = fmt.Errorf("TypeFood id: %s not found", id)
	}

	if !typeFood.IsNull() {
		db.Model(&typeFood).Where("id = ?", id).Updates(inputTypeFood)
	}

	return typeFood, err
}

func DeleteTypeFoodById(id interface{}) (typeFood models.TypeFood, isDelete bool,
	err error) {
	db := databases.Conn()
	db.First(&typeFood, id)

	if typeFood.IsNull() {
		err = fmt.Errorf("TypeFood id: %s not found", id)
	}

	if !typeFood.IsNull() {
		db.Delete(&typeFood, id)
		isDelete = true
	}

	return typeFood, isDelete, err
}
