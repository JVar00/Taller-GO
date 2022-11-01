package api

import (
	"github.com/labstack/echo/v4"
)

func (a *Api) RegisterRoutes(e *echo.Echo) {

	//handlers
	e.GET("/tasks", a.getTasks)

	e.GET("/tasks/:id", a.getTask)

}
