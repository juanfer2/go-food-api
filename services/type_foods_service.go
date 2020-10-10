package services

import (
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateTypeFoodService(inputTypeFood models.TypeFood) models.TypeFood {
	return repositories.CreateTypeFood(inputTypeFood)
}

func GetTypeFoodsService() []models.TypeFood {
	return repositories.GetTypeFoods()
}

func GetTypeFoodByIdService(id interface{}) (typeFood models.TypeFood, err error) {
	typeFood, err = repositories.GetTypeFoodById(id)
	return typeFood, err
}

func UpdateTypeFoodByIdService(id interface{}, inputTypeFood models.TypeFood) (typeFood models.TypeFood, err error) {
	typeFood, err = repositories.UpdateTypeFoodById(id, inputTypeFood)
	return typeFood, err
}

func DeleteTypeFoodByIdService(id interface{}) (typeFood models.TypeFood, isDelete bool, err error) {
	typeFood, isDelete, err = repositories.DeleteTypeFoodById(id)
	return typeFood, isDelete, err
}
