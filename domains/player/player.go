package player

import (
	"gorm.io/gorm"

	"github.com/hrz8/rahman-tennis/models"
)

// Repository is an interface of User domain for user model method
type (
	Repository interface {
		GetAll(db *gorm.DB) (*[]models.Player, error)
	}
)
