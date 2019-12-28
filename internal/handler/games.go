package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/handler/response"
	"github.com/robovarga/szlh-delegations/internal/repository"
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

func (h *GamesHandler) GetByListID(w http.ResponseWriter, r *http.Request) {
	listIDRaw := chi.URLParam(r, "listId")

	listID, err := strconv.Atoi(listIDRaw)
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	games, err := h.gameRepo.FindByListID(listID)
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	var responseRaw []*response.Game

	for _, game := range games {

		var refs []*response.Referee
		for _, ref := range game.Referees() {
			refs = append(refs, &response.Referee{
				RefID: ref.ID(),
				Name:  ref.Name(),
			})
		}

		responseRaw = append(
			responseRaw,
			&response.Game{
				GameUUID:   game.UUID().String(),
				ExternalID: game.ExternalID(),
				Home:       game.Home(),
				Away:       game.Away(),
				GameDate:   game.Date().String(),
				Referees:   refs,
			},
		)
	}

	response.Success(w, responseRaw)
}
