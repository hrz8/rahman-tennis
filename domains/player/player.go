package player

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	"github.com/hrz8/rahman-tennis/models"
)

type (
	// Repository is an interface of Player domain for user model method
	Repository interface {
		GetAll(db *gorm.DB) (*[]models.Player, error)
		Create(db *gorm.DB, p *models.Player) (*models.Player, error)
		GetByID(db *gorm.DB, id uuid.UUID) (*models.Player, error)
		UpdateOne(db *gorm.DB, p *models.Player, np *models.Player) (*models.Player, error)
	}

	// Usecase is an interface of Player domain for player sharable method
	Usecase interface {
		GetAll(c echo.Context) (*[]models.Player, error)
		Create(c echo.Context, p *models.Player) (*models.Player, error)
		GetByID(c echo.Context) (*models.Player, error)
		AddBall(c echo.Context) (*models.Player, error)
	}
)
