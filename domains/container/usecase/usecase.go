package repository

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	ContainerDomain "github.com/hrz8/rahman-tennis/domains/container"
	"github.com/hrz8/rahman-tennis/models"
	"github.com/hrz8/rahman-tennis/shared/lib"
)

type (
	handler struct {
		repository ContainerDomain.Repository
	}
)

// NewUsecase return implementation of methods in container-domain.Repository
func NewUsecase(repo ContainerDomain.Repository) ContainerDomain.Usecase {
	return &handler{
		repository: repo,
	}
}

func (h handler) GetAll(c echo.Context) (*[]models.Container, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	containers, err := h.repository.GetAll(db)
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (h handler) GetByPlayerID(c echo.Context) (*[]models.Container, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	paramsPlayerID := c.Param("playerID")
	if paramsPlayerID == "" {
		return nil, errors.New("invalid uuid parameter")
	}

	playerID, err := uuid.FromString(paramsPlayerID)
	if err != nil {
		return nil, errors.New("unkown parameter")
	}

	containers, err := h.repository.GetByPlayerID(db, playerID)
	if err != nil {
		return nil, err
	}

	return containers, nil
}
