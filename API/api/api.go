package api

import (
	"API/data"
	"API/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct{}

type TasksParams struct {
	Limit int `query:"limit"`
	//echo utiliza el parametro query para poder mapear estos keys values.
}

type TaskIdParams struct {
	Id int `param:"id"`
}

func (a *Api) getTasks(e echo.Context) error {

	params := &TasksParams{}

	err := e.Bind(params)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Query Params")
	}

	if params.Limit < 0 || params.Limit > len(data.TasksData) {
		return e.JSON(http.StatusBadRequest, "Invalid Query Params")
	}

	var to int

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(data.TasksData)
	}

	return e.JSON(http.StatusOK, data.TasksData[:to])
}

func (a *Api) postTask(e echo.Context) error {

	var task models.Task

	err := e.Bind(&task)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid parameters")
	}

	data.TasksData = append(data.TasksData, task)
	return e.NoContent(http.StatusCreated)
}

func (a *Api) deleteTask(e echo.Context) error {

	params := &TaskIdParams{}

	err := e.Bind(params)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid ID")
	}

	index := params.Id

	for i, task := range data.TasksData {
		if task.ID == index {
			// [a, b, c] esto concatena los los indices 0 y 2 y elimina el indice 1 (b)
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			return e.JSON(http.StatusOK, index)
		}
	}

	return e.JSON(http.StatusNotFound, "Unexisting ID")
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

//ESTO SE ELIMINA AL INICIO DE LA ACTIVIDAD, SE MANTIENE PARA REPASO PARA EXPLICAR EL FUNCIONAMIENTO DEL UPDATETASK

// func (a *Api) updateTask(e echo.Context) error {

// 	index, err := strconv.Atoi(e.Param("id"))

// 	var updatedTask models.Task

// 	if err != nil {
// 		return e.JSON(http.StatusBadRequest, "Invalid ID")
// 	}

// 	err = e.Bind(&updatedTask)
// 	if err != nil {
// 		return e.JSON(http.StatusBadRequest, "Invalid Content")
// 	}

// 	for i, task := range data.TasksData {
// 		if task.ID == index {
// 			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
// 			updatedTask.ID = task.ID
// 			data.TasksData = append(data.TasksData, updatedTask)
// 			return e.JSON(http.StatusOK, index)
// 		}
// 	}

// 	return e.JSON(http.StatusNotFound, "Unexisting ID")
// }
