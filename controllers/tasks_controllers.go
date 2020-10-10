package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/services"
	"github.com/juanfer2/api-rest-go/validations"
)

func parserBodyTask(c *fiber.Ctx) models.Task {
	newTaks := new(models.Task)
	c.BodyParser(newTaks)
	return *newTaks
}

// Create Task
func CreateTasks(c *fiber.Ctx) error {
	var inputTask models.Task
	inputTask = parserBodyTask(c)
	err := validations.ValidateCreateTask(inputTask)

	if err != nil {
		return c.Status(422).JSON(err)
	}
	data := services.CreateTaskService(inputTask)
	return c.JSON(data)
}

func GetTasks(c *fiber.Ctx) error {
	tasks := services.GetTasksService()
	return c.JSON(tasks)

}

func GetTasksById(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := services.GetTaskByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(task)
}

func UpdateTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	var inputTask models.Task
	inputTask = parserBodyTask(c)
	task, err := services.UpdateTaskByIdService(id, inputTask)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	return c.JSON(task)
}

func DeleteTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	msg := ""
	_, isDelete, err := services.DeleteTaskByIdService(id)

	if err != nil {
		msg := fmt.Sprint(err)
		return c.Status(404).JSON(msg)
	}

	if isDelete {
		msg = fmt.Sprintf("Task id %s has been deleted", id)
	} else {
		msg = fmt.Sprintf("Task id %s has not been deleted", id)
	}

	return c.JSON(msg)
}
