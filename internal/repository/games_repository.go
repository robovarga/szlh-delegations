package repository

import (
	"database/sql"

	"github.com/robovarga/szlh-delegations/internal/entity"
)

type GamesRepository struct {
	conn *sql.DB
}

func NewGamesRepository(conn *sql.DB) *GamesRepository {
	return &GamesRepository{conn}
}

func (repo *GamesRepository) Insert(game *entity.Game) error {
	query := `INSERT INTO games (home, away, external_id, venue, date) VALUES ($1, $2, $3, $4, $5);`

	_, err := repo.conn.Exec(query, game.Home(), game.Away(), game.ExternalID(), game.Venue(), game.Date())

	return err
}
