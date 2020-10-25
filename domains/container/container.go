package container

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	"github.com/hrz8/rahman-tennis/models"
)

type (
	// Repository is an interface of Container domain for user model method
	Repository interface {
		GetAll(db *gorm.DB) (*[]models.Container, error)
		GetByPlayerID(db *gorm.DB, pid uuid.UUID) (*[]models.Container, error)
		GetByID(db *gorm.DB, id uuid.UUID) (*models.Container, error)
		UpdateOne(db *gorm.DB, c *models.Container, nc *models.Container) (*models.Container, error)
	}

	// Usecase is an interface of Player domain for player sharable method
	Usecase interface {
		GetAll(c echo.Context) (*[]models.Container, error)
		GetByPlayerID(c echo.Context) (*[]models.Container, error)
	}
)
