package models

import (
	"github.com/gofrs/uuid"
)

// Player represent Player object for DB
type Player struct {
	ID          uuid.UUID   `gorm:"column:id;primaryKey" json:"id"`
	Name        string      `gorm:"column:name;size:255;not null" json:"name"`
	ReadyToPlay bool        `gorm:"column:ready_to_play;not null" json:"readyToPlay"`
	Containers  []Container `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"containers"`
}

// CreatePlayerPayload represent Player payload for creating Player
type CreatePlayerPayload struct {
	Name string `json:"name"`
}
