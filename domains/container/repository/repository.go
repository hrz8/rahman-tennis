package repository

import (
	"gorm.io/gorm"

	"github.com/gofrs/uuid"
	ContainerDomain "github.com/hrz8/rahman-tennis/domains/container"
	"github.com/hrz8/rahman-tennis/models"
)

type (
	handler struct {
		db *gorm.DB
	}
)

// NewRepository return implementation of methods in transaction.Repositoru
func NewRepository(db *gorm.DB) ContainerDomain.Repository {
	return &handler{
		db: db,
	}
}

func (h handler) GetAll(db *gorm.DB) (*[]models.Container, error) {
	containers := &[]models.Container{}
	if err := db.Model(&models.Container{}).Find(containers).Error; err != nil {
		return containers, err
	}
	return containers, nil
}

func (h handler) GetByPlayerID(db *gorm.DB, pid uuid.UUID) (*[]models.Container, error) {
	containers := &[]models.Container{}
	if err := db.Where(&models.Container{
		PlayerID: pid,
	}).Find(containers).Error; err != nil {
		return nil, err
	}
	return containers, nil
}

func (h handler) GetByID(db *gorm.DB, id uuid.UUID) (*models.Container, error) {
	c := models.Container{}
	if err := db.Take(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (h handler) UpdateOne(db *gorm.DB, c *models.Container, nc *models.Container) (*models.Container, error) {
	if err := db.Model(&c).Updates(nc).Error; err != nil {
		return nil, err
	}
	return nc, nil
}
