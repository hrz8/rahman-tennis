package repository

import (
	"gorm.io/gorm"

	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
	"github.com/hrz8/rahman-tennis/models"
)

type (
	handler struct {
		db *gorm.DB
	}
)

// NewRepository return implementation of methods in transaction.Repositoru
func NewRepository(db *gorm.DB) PlayerDomain.Repository {
	return &handler{
		db: db,
	}
}

func (h handler) GetAll(db *gorm.DB) (*[]models.Player, error) {
	players := &[]models.Player{}
	if err := db.Model(&models.Player{}).Find(players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (h handler) Create(db *gorm.DB, u *models.Player) (*models.Player, error) {
	if err := db.Debug().Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
