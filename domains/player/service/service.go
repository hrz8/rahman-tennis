package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
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
