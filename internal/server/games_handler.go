package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/robovarga/szlh-delegations/internal/repository"
	"github.com/sirupsen/logrus"
)

type GamesHandler struct {
	gameRepo *repository.GamesRepository
	logger   *logrus.Logger
}

func NewGamesHandler(gamesRepository *repository.GamesRepository,
	logger *logrus.Logger) *GamesHandler {
	return &GamesHandler{
		gameRepo: gamesRepository,
		logger:   logger,
	}
}

func (h *GamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	listIDRaw := chi.URLParam(r, "id")

	listID, err := strconv.Atoi(listIDRaw)
	if err != nil {
		h.logger.Error(err)
	}

	games, err := h.gameRepo.FindByListID(listID)
	if err != nil {
		h.logger.Error(err)
	}

	var responseRaw []gameResponse

	for _, game := range games {
		responseRaw = append(
			responseRaw,
			gameResponse{
				GameID:     game.ID().String(),
				ExternalID: game.ExternalID(),
				Home:       game.Home(),
				Away:       game.Away(),
				GameDate:   game.Date().String(),
			},
		)
	}

	response, err := json.Marshal(responseRaw)
	if err != nil {
		h.logger.Error(err)
	}

	w.WriteHeader(200)
	_, _ = w.Write(response)
}

type gameResponse struct {
	GameID     string `json:"game_id"`
	ExternalID int    `json:"external_id"`
	Home       string `json:"home"`
	Away       string `json:"away"`
	GameDate   string `json:"game_date"`
}
