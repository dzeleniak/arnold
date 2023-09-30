package controllers

import (
	"net/http"

	"github.com/dzeleniak/arnold/services"
	"github.com/labstack/echo/v4"
)

type Controllers struct {
	MovementController
	MovementSetController
}

func New(s *services.Services) *Controllers {
	return &Controllers{
		MovementController: &movementController{s.Movement},
		MovementSetController: &movementSetController{s.MovementSet},
	}
}

func Echo() *echo.Echo {
	e := echo.New()

	return e;
}	
func SetDefault(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Success!");
	})
}

func SetApi(e *echo.Echo, c *Controllers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	// g.Use(m)

	g.GET("/", func (ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "API V1")
	})

	g.GET("/movements", c.MovementController.GetMovements)
	g.POST("/movements", c.MovementController.CreateMovement)
	g.PUT("/movements", c.MovementController.UpdateMovementById)
	g.DELETE("/movements/:id", c.MovementController.DeleteMovementById)

	g.GET("/sets", c.MovementSetController.GetMovementSets);
	g.POST("/sets", c.MovementSetController.CreateMovementSet);
	g.PUT("/sets", c.MovementSetController.UpdateMovementSet);
	g.DELETE("/sets/:id", c.MovementSetController.DeleteMovementSet);
}