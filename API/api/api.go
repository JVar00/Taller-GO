package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct{}

type TaskIdParams struct {
	Id int `param:"id"`
}

//funcion getTasks

func (a *Api) getTask(e echo.Context) error {

	//Uso del PARAMS

	//Busqueda

	return e.JSON(http.StatusNotFound, "Unexisting ID")
}
