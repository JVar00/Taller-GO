package main

import (
	"API/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", handleIndex)

	a := &api.Api{}
	a.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

}

func handleIndex(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
}
