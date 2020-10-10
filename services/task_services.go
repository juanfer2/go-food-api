package services

import (
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/repositories"
)

func CreateTaskService(inputTaks models.Task) models.Task {
	return repositories.CreateTaks(inputTaks)
}

func GetTasksService() []models.Task {
	return repositories.GetTaks()
}

func GetTaskByIdService(id interface{}) (task models.Task, err error) {
	task, err = repositories.GetTasksById(id)
	return task, err
}

func UpdateTaskByIdService(id interface{}, inputTask models.Task) (task models.Task, err error) {
	task, err = repositories.UpdateTaskById(id, inputTask)
	return task, err
}

func DeleteTaskByIdService(id interface{}) (task models.Task, isDelete bool, err error) {
	task, isDelete, err = repositories.DeleteTaskById(id)
	return task, isDelete, err
}
