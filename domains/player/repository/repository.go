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
	var err error
	players := &[]models.Player{}
	err = db.Model(&models.Player{}).Find(players).Error
	if err != nil {
		return &[]models.Player{}, err
	}
	return players, err
}
