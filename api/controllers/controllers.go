package controllers

import (
	"github.com/dzeleniak/arnold/services"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type Controllers struct {
	MovementController
}

func New(s *services.Services) *Controllers {
	return &Controllers{
		&movementController{s.Movement},
	}
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger());
	e.Use(middleware.Recover());
	e.Pre(middleware.RemoveTrailingSlash());
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowedOrigins: []string{"*"},
	}))

	return e;
}	

func SetApi(e *echo.Echo, c *Controllers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	g.Use(m)

	g.GET("/movements", c.MovementController.GetMovements)
	g.POST("/movements", c.MovementController.CreateMovement)
	g.PUT("/movements", c.MovementController.UpdateMovementById)
	g.DELETE("/movements/:id", c.MovementController.DeleteMovementById)
}