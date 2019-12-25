package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

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
		h.logger.Error(err)
	}

	var responseRaw []listsResponse

	for _, list := range lists {
		responseRaw = append(
			responseRaw,
			listsResponse{
				ID:   list.ListID(),
				Name: list.Name(),
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

type listsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
