package container

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	"github.com/hrz8/rahman-tennis/models"
)

// Repository is an interface of User domain for user model method
type (
	Repository interface {
		GetAll(db *gorm.DB) (*[]models.Container, error)
		GetByID(db *gorm.DB, id uuid.UUID) (*models.Container, error)
		UpdateOne(db *gorm.DB, c *models.Container, nc *models.Container) (*models.Container, error)
	}
)
