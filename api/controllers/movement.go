package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dzeleniak/arnold/models"
	"github.com/dzeleniak/arnold/services"
	"github.com/labstack/echo/v4"
)

type (
	MovementController  interface {
		GetMovements(ctx echo.Context) error
		CreateMovement(ctx echo.Context) error
		UpdateMovementById(ctx echo.Context) error
		DeleteMovementById(ctx echo.Context) error
	}

	movementController struct {
		services.MovementService
	}
)

func (c *movementController) GetMovements(ctx echo.Context) error {
	movements, err := c.MovementService.GetMovements()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,  err.Error())
	}

	return ctx.JSON(http.StatusOK, movements);
}

func (c *movementController) CreateMovement(ctx echo.Context) error {
	var m *models.Movement

	if err := ctx.Bind(&m); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := c.MovementService.CreateMovement(m);

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (c *movementController) UpdateMovementById(ctx echo.Context) error {
	var m *models.Movement

	if err := ctx.Bind(&m); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error());
	}

	id, err := c.MovementService.UpdateMovementById(m);

	if err == sql.ErrNoRows {
		return ctx.JSON(http.StatusNotFound, err.Error());
	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error());
	}

	return ctx.JSON(http.StatusOK, id)
}

func (c *movementController) DeleteMovementById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"));

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error());
	}

	err = c.MovementService.DeleteMovement(id);

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "OK");
}