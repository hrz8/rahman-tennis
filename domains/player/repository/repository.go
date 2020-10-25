package repository

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"

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
	if err := db.Preload("Containers").Find(players).Error; err != nil {
		return players, err
	}
	return players, nil
}

func (h handler) Create(db *gorm.DB, p *models.Player) (*models.Player, error) {
	if err := db.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (h handler) GetByID(db *gorm.DB, id uuid.UUID) (*models.Player, error) {
	p := models.Player{}
	if err := db.Preload("Containers").Take(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (h handler) UpdateOne(db *gorm.DB, p *models.Player, np *models.Player) (*models.Player, error) {
	if err := db.Model(&p).Updates(np).Error; err != nil {
		return nil, err
	}
	return np, nil
}
