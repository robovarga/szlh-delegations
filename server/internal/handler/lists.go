package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/robovarga/szlh-delegations/internal/handler/response"
	"github.com/robovarga/szlh-delegations/internal/repository"
)

type ListsHandler struct {
	listRepo *repository.ListRepository
	logger   *logrus.Logger
}

func NewListsHandler(listRepository *repository.ListRepository,
	logger *logrus.Logger) *ListsHandler {
	return &ListsHandler{
		listRepo: listRepository,
		logger:   logger,
	}
}

func (h *ListsHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	lists, err := h.listRepo.GetLists()
	if err != nil {
		response.InternalError(w, h.logger, err)
		return
	}

	var responseRaw []*response.List

	for _, list := range lists {
		responseRaw = append(
			responseRaw,
			&response.List{
				ID:   list.ListID(),
				Name: list.Name(),
			},
		)
	}

	response.Success(w, responseRaw)
}
