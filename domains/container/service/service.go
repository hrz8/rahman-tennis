package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	ContainerDomain "github.com/hrz8/rahman-tennis/domains/container"
	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
)

type (
	handler struct {
		usecase       ContainerDomain.Usecase
		playerUsecase PlayerDomain.Usecase
	}
)

// InitService will return REST of container-domain
func InitService(e *echo.Echo, usecase ContainerDomain.Usecase, playerUsecase PlayerDomain.Usecase) {
	h := handler{
		usecase:       usecase,
		playerUsecase: playerUsecase,
	}

	e.GET("/api/v1/containers", h.GetAll)
	e.GET("/api/v1/containers/player/:playerID", h.GetByPlayerID)
}

func (h handler) GetAll(c echo.Context) error {
	containers, err := h.usecase.GetAll(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   containers,
	})
}

func (h handler) GetByPlayerID(c echo.Context) error {
	player, err := h.playerUsecase.GetByID(c)
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

	containers, err := h.usecase.GetByPlayerID(c)
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
		"data": echo.Map{
			"isVerified": player.ReadyToPlay,
			"containers": containers,
		},
	})
}
