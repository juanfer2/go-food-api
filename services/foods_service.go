package services

import (
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateFoodService(inputFood models.Food) models.Food {
	return repositories.CreateFood(inputFood)
}

func GetFoodsService() []models.Food {
	return repositories.GetFoods()
}

func GetFoodByIdService(id interface{}) (food models.Food, err error) {
	food, err = repositories.GetFoodById(id)
	return food, err
}

func UpdateFoodByIdService(id interface{}, inputFood models.Food) (food models.Food, err error) {
	food, err = repositories.UpdateFoodById(id, inputFood)
	return food, err
}

func DeleteFoodByIdService(id interface{}) (food models.Food, isDelete bool, err error) {
	food, isDelete, err = repositories.DeleteFoodById(id)
	return food, isDelete, err
}
