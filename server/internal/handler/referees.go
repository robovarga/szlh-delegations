package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/robovarga/szlh-delegations/internal/handler/response"
	"github.com/robovarga/szlh-delegations/internal/repository"
	"github.com/sirupsen/logrus"
)

type RefereesHandler struct {
	refsRepo  *repository.RefereesRepository
	gamesRepo *repository.GamesRepository
	logger    *logrus.Logger
}

func NewRefereesHandler(refsRepo *repository.RefereesRepository,
	gamesRepo *repository.GamesRepository,
	logger *logrus.Logger) *RefereesHandler {
	return &RefereesHandler{
		refsRepo:  refsRepo,
		gamesRepo: gamesRepo,
		logger:    logger,
	}
}

func (h *RefereesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	refs, err := h.refsRepo.GetAll()
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	var responseRaw []*response.Referee

	for _, ref := range refs {
		responseRaw = append(
			responseRaw,
			&response.Referee{
				RefID:      ref.ID(),
				Name:       ref.Name(),
				DateAdd:    ref.DateAdd(),
				DateUpdate: ref.DateUpdate(),
			},
		)
	}

	response.Success(w, responseRaw)
}

func (h *RefereesHandler) GetReferee(w http.ResponseWriter, r *http.Request) {
	refIDRaw := chi.URLParam(r, "refId")
	refID, err := strconv.Atoi(refIDRaw)
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	games, err := h.gamesRepo.FindByRefId(refID)
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	var responseRaw []*response.Game

	for _, game := range games {
		responseRaw = append(
			responseRaw,
			&response.Game{
				GameUUID:   game.UUID().String(),
				ExternalID: game.ExternalID(),
				Home:       game.Home(),
				Away:       game.Away(),
				GameDate:   game.Date().String(),
			},
		)
	}

	response.Success(w, responseRaw)
}
