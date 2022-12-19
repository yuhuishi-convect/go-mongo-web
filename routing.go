package main

import (
	"github.com/labstack/echo/v4"
)

// routing the url paths to the handler functions
func ConfigureRouting(e *echo.Echo, h *TodoItemHandler) {
	e.GET("/todo", h.FindAll)
	e.GET("/todo/:id", h.FindByID)
	e.POST("/todo", h.Save)
}
