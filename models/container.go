package models

import (
	"github.com/gofrs/uuid"
)

// Container represents a Container entity object
type Container struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	PlayerID uuid.UUID `gorm:"column:player_id;not null;index" json:"playerId"`
	Capacity uint64    `gorm:"column:capacity;min:1;default:1" json:"capacity"`
	BallQty  uint64    `gorm:"column:ball_qty;min:0;default:0" json:"ballQty"`
	Player   Player    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"player"`
}

// CreateContainerPayload represent Container payload for creating Container
type CreateContainerPayload struct {
	PlayerID uuid.UUID `json:"playerId"`
	Capacity int       `json:"capacity"`
}
