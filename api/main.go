package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET: /");
	})

	e.POST("/", func(c echo.Context) error { 
		return c.String(http.StatusOK, "POST: /");
	})

	e.PUT("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT: /");
	})

	e.DELETE("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Delete: /")
	})

	e.GET("/Movements", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET: /Movements");
	})

	e.POST("/Movements", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST: /Movements");
	})

	e.PUT("/Movements", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT: /Movements")
	})

	e.DELETE("/Movements", func(c echo.Context) error { 
		return c.String(http.StatusOK, "DELETE: /Movements")
	})

	e.Logger.Fatal(e.Start(":1323"));
}