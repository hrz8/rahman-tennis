package service

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
	"github.com/hrz8/rahman-tennis/models"
	"github.com/hrz8/rahman-tennis/shared/lib"
)

type (
	handler struct {
		repository PlayerDomain.Repository
	}
)

// InitService will return REST
func InitService(e *echo.Echo, repo PlayerDomain.Repository) {
	h := handler{
		repository: repo,
	}

	e.GET("/api/v1/players", h.GetAll)
	e.POST("/api/v1/players", h.Create)
}

func (h handler) GetAll(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	players, err := h.repository.GetAll(db)
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
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	payload := &models.CreatePlayerPayload{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
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
	player := &models.Player{
		ID:          id,
		Name:        payload.Name,
		ReadyToPlay: false,
		Containers:  containers,
	}

	playerResponse, err := h.repository.Create(db, player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   playerResponse,
	})
}
