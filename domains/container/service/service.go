package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	ContainerDomain "github.com/hrz8/rahman-tennis/domains/container"
	"github.com/hrz8/rahman-tennis/shared/lib"
)

type (
	handler struct {
		repository ContainerDomain.Repository
	}
)

// InitService will return REST
func InitService(e *echo.Echo, repo ContainerDomain.Repository) {
	h := handler{
		repository: repo,
	}

	e.GET("/api/v1/containers", h.GetAll)
}

func (h handler) GetAll(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	containers, err := h.repository.GetAll(db)
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
