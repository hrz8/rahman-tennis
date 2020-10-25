package repository

import (
	"errors"
	"math/rand"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	ContainerDomain "github.com/hrz8/rahman-tennis/domains/container"
	PlayerDomain "github.com/hrz8/rahman-tennis/domains/player"
	"github.com/hrz8/rahman-tennis/models"
	"github.com/hrz8/rahman-tennis/shared/lib"
)

type (
	handler struct {
		repository          PlayerDomain.Repository
		containerRepository ContainerDomain.Repository
	}
)

// NewUsecase return implementation of methods in transaction.Repositoru
func NewUsecase(repo PlayerDomain.Repository, containerRepo ContainerDomain.Repository) PlayerDomain.Usecase {
	return &handler{
		repository:          repo,
		containerRepository: containerRepo,
	}
}

func (h handler) GetAll(c echo.Context) (*[]models.Player, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	players, err := h.repository.GetAll(db)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (h handler) Create(c echo.Context, p *models.Player) (*models.Player, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	playerResponse, err := h.repository.Create(db, p)
	if err != nil {
		return nil, err
	}

	return playerResponse, nil
}

func (h handler) GetByID(c echo.Context) (*models.Player, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	paramsPlayerID := c.Param("playerID")
	if paramsPlayerID == "" {
		return nil, errors.New("invalid uuid parameter")
	}

	playerID, err := uuid.FromString(paramsPlayerID)
	if err != nil {
		return nil, errors.New("unkown parameter")
	}

	player, err := h.repository.GetByID(db, playerID)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (h handler) AddBall(c echo.Context) (*models.Player, error) {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	player, err := h.GetByID(c)
	if err != nil {
		return nil, err
	}

	if player.ReadyToPlay {
		return nil, errors.New("user already ready")
	}

	containerIdx := rand.Intn(len((*player).Containers))
	randomContainer := (*player).Containers[containerIdx]

	updatedContainer, err := h.containerRepository.UpdateOne(db, &randomContainer, &models.Container{
		BallQty: randomContainer.BallQty + 1,
	})
	if err != nil {
		return nil, err
	}

	randomContainer.BallQty = updatedContainer.BallQty
	if randomContainer.BallQty == randomContainer.Capacity {
		h.repository.UpdateOne(db, player, &models.Player{
			ReadyToPlay: true,
		})
	}

	(*player).Containers[containerIdx] = randomContainer
	return player, nil
}
