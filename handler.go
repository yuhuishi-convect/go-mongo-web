package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// handler functions for requests

// TodoItemHandler is the handler for todo items
type TodoItemHandler struct {
	repository TodoItemRepository
}

// NewTodoItemHandler creates a new TodoItemHandler
func NewTodoItemHandler(repository TodoItemRepository) *TodoItemHandler {
	return &TodoItemHandler{repository: repository}
}

// FindAll returns all todo items
func (h *TodoItemHandler) FindAll(c echo.Context) error {
	items, err := h.repository.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(200, items)
}

// FindByID returns a todo item by id
func (h *TodoItemHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	item, err := h.repository.FindByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(200, item)
}

// Save saves a todo item
func (h *TodoItemHandler) Save(c echo.Context) error {
	var item TodoItem
	err := c.Bind(&item)
	if err != nil {
		return err
	}
	item.ID = primitive.NewObjectID()
	err = h.repository.Save(item)
	if err != nil {
		return err
	}
	return c.JSON(201, item)
}
