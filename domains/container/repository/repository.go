package repository

import (
	"gorm.io/gorm"

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
	var err error
	containers := &[]models.Container{}
	err = db.Model(&models.Container{}).Find(containers).Error
	if err != nil {
		return &[]models.Container{}, err
	}
	return containers, err
}
