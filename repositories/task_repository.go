package repositories

import (
	"fmt"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
)

func GetTaks() []models.Task {
	db := databases.Conn()
	var tasks []models.Task
	db.Find(&tasks)
	return tasks
}

func GetTasksById(id interface{}) (task models.Task, err error) {
	db := databases.Conn()
	db.First(&task, id)

	if task.IsNull() {
		err = fmt.Errorf("Task id: %s not found", id)
	}
	return task, err
}

func CreateTaks(inputTask models.Task) models.Task {
	db := databases.Conn()
	task := inputTask
	db.Create(&task)

	return task
}

func UpdateTaskById(id interface{}, inputTask models.Task) (task models.Task, err error) {
	db := databases.Conn()
	db.First(&task, id)

	if task.IsNull() {
		err = fmt.Errorf("Task id: %s not found", id)
	}

	if !task.IsNull() {
		db.Model(&task).Where("id = ?", id).Updates(inputTask)
	}

	return task, err
}

func DeleteTaskById(id interface{}) (task models.Task, isDelete bool, err error) {
	db := databases.Conn()
	db.First(&task, id)

	if task.IsNull() {
		err = fmt.Errorf("Task id: %s not found", id)
	}

	if !task.IsNull() {
		db.Delete(&task, id)
		isDelete = true
	}

	return task, isDelete, err
}
