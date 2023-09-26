package controllers

import (
	"github.com/dzeleniak/arnold/services"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	MovementController
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		&movementController{s.Movement},
	}
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger());
	e.Use(middleware.Recover());
	e.Pre(middleware.RemoveTrailingSlash);
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowedOrigins: []string{"*"},
	}))

	return e;
}	