package controllers

import (
	"net/http"
	"strconv"

	"github.com/dzeleniak/arnold/models"
	"github.com/dzeleniak/arnold/services"
	"github.com/labstack/echo/v4"
)

type (
	MovementSetController interface {
		GetMovementSets(ctx echo.Context) error
		CreateMovementSet(ctx echo.Context) error
		UpdateMovementSet(ctx echo.Context) error
		DeleteMovementSet(ctx echo.Context) error
		GetMovementSetsByMovementID(ctx echo.Context) error
	}	

	movementSetController struct {
		services.MovementSetService
	}
)

func (c *movementSetController) GetMovementSets(ctx echo.Context) error {
	var sets []models.MovementSet
	sets, err := c.MovementSetService.GetMovementSets()
		
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, sets)
}

func (c *movementSetController) CreateMovementSet(ctx echo.Context) error {
	var set *models.MovementSet

	err := ctx.Bind(&set)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())	
	}

	id, err := c.MovementSetService.CreateMovementSet(set)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (c *movementSetController) UpdateMovementSet(ctx echo.Context) error {

	var set *models.MovementSet

	if	err := ctx.Bind(&set); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := c.MovementSetService.UpdateMovementSet(set);

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id);
}

func (c *movementSetController) DeleteMovementSet(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"));

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.MovementSetService.DeleteMovementSet(int64(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (c *movementSetController) GetMovementSetsByMovementID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"));

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	sets, err := c.MovementSetService.GetMovementSetsByMovementID(int64(id))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, sets)
}