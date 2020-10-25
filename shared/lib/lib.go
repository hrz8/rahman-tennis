package lib

import (
	"math/rand"

	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	PlayerRepository "github.com/hrz8/rahman-tennis/domains/player/repository"
	"github.com/hrz8/rahman-tennis/models"
)

type (
	// AppContext return custom application context
	AppContext struct {
		echo.Context
		MysqlSession *gorm.DB
	}

	// CustomValidator return custom request validator
	CustomValidator struct {
		Validator *validator.Validate
	}
)

// Validate will validate given input with related struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// Migrate will doing migration for DB
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Player{},
		&models.Container{},
	)
	migrationRepo := PlayerRepository.NewRepository(db)

	rahmanContainers := make([]models.Container, 0)
	for i := 0; i < 7; i++ {
		cid, _ := uuid.NewV4()
		rahmanContainers = append(rahmanContainers, models.Container{
			ID:       cid,
			Capacity: 7,
			BallQty:  0,
		})
	}
	migrationRepo.Create(db, &models.Player{
		ID:          uuid.FromStringOrNil("4db77dd4-09b0-4633-aed2-a8382e17a748"),
		Name:        "rahman",
		ReadyToPlay: false,
		Containers:  rahmanContainers,
	})

	verifiedContainers := make([]models.Container, 0)
	vid, _ := uuid.NewV4()
	verifiedContainers = append(verifiedContainers, models.Container{
		ID:       vid,
		Capacity: 7,
		BallQty:  7,
	})
	for i := 0; i < 6; i++ {
		cid, _ := uuid.NewV4()
		verifiedContainers = append(verifiedContainers, models.Container{
			ID:       cid,
			Capacity: 7,
			BallQty:  uint64(rand.Intn(7)),
		})
	}
	migrationRepo.Create(db, &models.Player{
		ID:          uuid.FromStringOrNil("b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755"),
		Name:        "verified player",
		ReadyToPlay: true,
		Containers:  verifiedContainers,
	})
}
