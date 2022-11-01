package api

import (
	"API/data"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct{}

type TaskIdParams struct {
	Id int `param:"id"`
}

func (a *Api) getTasks(e echo.Context) error {
	return e.JSON(http.StatusOK, data.TasksData)
}

func (a *Api) getTask(e echo.Context) error {

	params := &TaskIdParams{}

	err := e.Bind(params)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid ID")
	}

	index := params.Id

	for _, task := range data.TasksData {
		if task.ID == index {
			return e.JSON(http.StatusOK, task)
		}
	}

	return e.JSON(http.StatusNotFound, "Unexisting ID")
}
