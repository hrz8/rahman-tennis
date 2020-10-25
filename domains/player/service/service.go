package service

import (
	"math/rand"
	"net/http"

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

// InitService will return REST
func InitService(e *echo.Echo, repo PlayerDomain.Repository, containerRepo ContainerDomain.Repository) {
	h := handler{
		repository:          repo,
		containerRepository: containerRepo,
	}

	e.GET("/api/v1/players", h.GetAll)
	e.POST("/api/v1/players", h.Create)
	e.GET("/api/v1/players/:playerID", h.GetByID)
	e.PUT("/api/v1/players/:playerID", h.AddBall)
}

func (h handler) GetAll(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	players, err := h.repository.GetAll(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   players,
	})
}

func (h handler) Create(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	payload := &models.CreatePlayerPayload{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	containers := make([]models.Container, 0)
	for i := 0; i < int(payload.ContainerQty); i++ {
		cid, _ := uuid.NewV4()
		containers = append(containers, models.Container{
			ID:       cid,
			Capacity: payload.ContainerCapacity,
			BallQty:  0,
		})
	}

	id, _ := uuid.NewV4()
	player := &models.Player{
		ID:          id,
		Name:        payload.Name,
		ReadyToPlay: false,
		Containers:  containers,
	}

	playerResponse, err := h.repository.Create(db, player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   playerResponse,
	})
}

func (h handler) GetByID(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	paramsPlayerID := c.Param("playerID")
	if paramsPlayerID == "" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  "invalid uuid parameter",
		})
	}

	playerID, err := uuid.FromString(paramsPlayerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  "unkown parameter",
		})
	}

	player, err := h.repository.GetByID(db, playerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   player,
	})
}

func (h handler) AddBall(c echo.Context) error {
	ac := c.(*lib.AppContext)
	db := ac.MysqlSession

	paramsPlayerID := c.Param("playerID")
	if paramsPlayerID == "" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  "invalid uuid parameter",
		})
	}

	playerID, err := uuid.FromString(paramsPlayerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  "unkown parameter",
		})
	}

	player, err := h.repository.GetByID(db, playerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	containerIdx := rand.Intn(len((*player).Containers))
	randomContainer := (*player).Containers[containerIdx]
	newContainer := models.Container{
		BallQty: randomContainer.BallQty + 1,
	}

	updatedContainer, err := h.containerRepository.UpdateOne(db, &randomContainer, &newContainer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	randomContainer.BallQty = updatedContainer.BallQty
	(*player).Containers[containerIdx] = randomContainer
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   *player,
	})
}
