package service

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
	"github.com/hrz8/rahman-tennis/models"
)

type (
	handler struct {
		usecase PlayerDomain.Usecase
	}
)

// InitService will return REST
func InitService(e *echo.Echo, usecase PlayerDomain.Usecase) {
	h := handler{
		usecase: usecase,
	}

	e.GET("/api/v1/players", h.GetAll)
	e.POST("/api/v1/players", h.Create)
	e.GET("/api/v1/players/:playerID", h.GetByID)
	e.PUT("/api/v1/players/:playerID", h.AddBall)
}

func (h handler) GetAll(c echo.Context) error {
	players, err := h.usecase.GetAll(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   players,
	})
}

func (h handler) Create(c echo.Context) error {
	payload := &models.CreatePlayerPayload{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	containers := make([]models.Container, 0)
	for i := 0; i < int(payload.ContainerQty); i++ {
		cid, _ := uuid.NewV4()
		containers = append(containers, models.Container{
			ID:       cid,
			Capacity: payload.ContainerCapacity,
			BallQty:  0,
		})
	}

	id, _ := uuid.NewV4()
	newPlayer := &models.Player{
		ID:          id,
		Name:        payload.Name,
		ReadyToPlay: false,
		Containers:  containers,
	}

	player, err := h.usecase.Create(c, newPlayer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   player,
	})
}

func (h handler) GetByID(c echo.Context) error {
	player, err := h.usecase.GetByID(c)
	if err != nil {
		var status int
		switch err.Error() {
		case "invalid uuid parameter", "unkown parameter":
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, echo.Map{
			"status": status,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   player,
	})
}

func (h handler) AddBall(c echo.Context) error {
	player, err := h.usecase.AddBall(c)
	if err != nil {
		var status int
		switch err.Error() {
		case "user already ready":
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, echo.Map{
			"status": status,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   player,
	})
}
