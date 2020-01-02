package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Referee struct {
	RefID      int       `json:"id"`
	Name       string    `json:"name"`
	DateAdd    time.Time `json:"date_add"`
	DateUpdate time.Time `json:"date_update"`
}

type Game struct {
	GameUUID   string     `json:"game_uuid"`
	ExternalID int        `json:"external_id"`
	Home       string     `json:"home"`
	Away       string     `json:"away"`
	GameDate   string     `json:"game_date"`
	Referees   []*Referee `json:"referees,omitempty"`
}

type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(body)
}

func InternalError(w http.ResponseWriter, log *logrus.Logger, err error) {
	errRes := Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	res, err := json.Marshal(errRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write(res)
}
