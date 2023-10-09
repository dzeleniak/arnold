package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dzeleniak/arnold/controllers"
	"github.com/dzeleniak/arnold/models"
	"github.com/dzeleniak/arnold/services"
	"github.com/stretchr/testify/assert"
)

type MockMovementService struct {
	services.MovementService
	MockGetMovements func() ([]models.Movement, error)
	MockDeleteMovementById func(id int) error
	MockCreateMovement func(m *models.Movement) (int64, error)
}

func (m *MockMovementService) GetMovements() ([]models.Movement, error) {
	return m.MockGetMovements()
}

func (m *MockMovementService) DeleteMovement(id int) error {
	return m.MockDeleteMovementById(id)
}

func (m *MockMovementService) CreateMovement(movement *models.Movement) (int64, error) {
	return m.MockCreateMovement(movement)
}

func TestCreateMovementSuccessCase(t *testing.T) {
	s := &MockMovementService{
		MockCreateMovement: func(m *models.Movement) (int64, error) {
			return 1, nil;
		},
	}

	mockService := &services.Services{ Movement: s}

	e := controllers.Echo()
	controllers := controllers.New(mockService)
	
	bodyObj := map[string]string{
		"name" : "asdfasdf",
	}

	body, _ := json.Marshal(bodyObj);
	
	req := httptest.NewRequest(http.MethodPost, "/api/v1/movements", bytes.NewReader(body))
	req.Header.Add("Content-type", "application/json")
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	assert.NoError(t, controllers.MovementController.CreateMovement(c))
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetMovementsSuccessCase(t *testing.T) {
	// Create instance of Mock Service with Get function set to return positive result
	s := &MockMovementService{
		MockGetMovements: func() ([]models.Movement, error) {
			var r []models.Movement
			return r, nil
		},
	}

	// Use Mock Service to create an instance of our services parent class
	mockService := &services.Services{Movement: s}

	// Create new Echo instance and controllers
	e := controllers.Echo();
	controllers := controllers.New(mockService)

	// Outline request and response
	req := httptest.NewRequest(http.MethodGet, "/api/v1/movement", nil)
	res := httptest.NewRecorder()

	// Create context to pass to the controller
	c := e.NewContext(req, res);

	// Check success case
	assert.NoError(t, controllers.MovementController.GetMovements(c))
}

func TestGetMovements500Case(t *testing.T) {
	// Create Mock Movement Service instance that returns an error
	s := &MockMovementService{
		MockGetMovements: func() ([]models.Movement, error) {
			return nil, errors.New("WE HAVE NO SETS :(((")
		},
	}

	// Create instance of parent class using our mock class
	mockService := &services.Services{ Movement: s }

	e := controllers.Echo()
	controllers := controllers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v2/movements", nil);
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	assert.NoError(t, controllers.GetMovements(c))
	assert.Equal(t, http.StatusInternalServerError, res.Code)
}

func TestDeleteMovementSuccessCase(t *testing.T) {
	s := &MockMovementService{
		MockDeleteMovementById: func(id int) error {
			return nil;
		},
	}

	mockService := &services.Services{ Movement: s }

	e := controllers.Echo()
	controllers := controllers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	res := httptest.NewRecorder();

	c := e.NewContext(req, res)

	c.SetPath("api/v1/movements/:id")
	c.SetParamNames("id")
	c.SetParamValues("1");

	assert.NoError(t, controllers.MovementController.DeleteMovementById(c));
	assert.Equal(t, http.StatusOK, res.Code);
}

func TestDeleteMovement400Case(t *testing.T) {
	s := &MockMovementService{
		MockDeleteMovementById: func(id int) error {
			return nil;
		},
	}

	mockService := &services.Services{ Movement: s }

	e := controllers.Echo();
	controllers := controllers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil);
	res := httptest.NewRecorder();

	c := e.NewContext(req, res);
	c.SetPath("api/v1/movements/:id")

	assert.NoError(t, controllers.DeleteMovementById(c))
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestDeleteMovement404Case(t *testing.T) {
	s := &MockMovementService{
		MockDeleteMovementById: func(id int) error {
			return sql.ErrNoRows
		},
	}

	mockService := &services.Services{Movement: s}

	e := controllers.Echo()
	controllers := controllers.New(mockService);

	req := httptest.NewRequest(http.MethodDelete, "/", nil);
	res := httptest.NewRecorder();

	c := e.NewContext(req,res)
	c.SetPath("api/v1/movements/:id")
	c.SetParamNames("id")
	c.SetParamValues("54354521")

	assert.NoError(t, controllers.MovementController.DeleteMovementById(c))
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestDeleteMovement500Case(t *testing.T) {
	s := &MockMovementService{
		MockDeleteMovementById: func(id int) error {
			return errors.New("fml :(")
		},
	}

	mockService := &services.Services{
		Movement: s,
	}

	e := controllers.Echo();
	controllers := controllers.New(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req,res)
	c.SetPath("api/v1/movements/:id")
	c.SetParamNames("id")
	c.SetParamValues("12")

	assert.NoError(t, controllers.MovementController.DeleteMovementById(c))
	assert.Equal(t, http.StatusInternalServerError, res.Code)
}