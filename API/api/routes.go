package api

import (
	"github.com/labstack/echo/v4"
)

func (a *Api) RegisterRoutes(e *echo.Echo) {

	//handlers
	e.GET("/tasks", a.getTasks)

	e.GET("/tasks/:id", a.getTask)
	e.DELETE("/tasks/:id", a.deleteTask)
	//el :id es un parametro que recibimos

	e.POST("/tasks", a.postTask)

}

//ESTO SE ELIMINA AL INICIO DE LA ACTIVIDAD
//e.PUT("/tasks/:id", a.updateTask)
