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
	query := `INSERT INTO games (game_id, home_team, away_team, external_id, list_id, venue, game_date)
		VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := repo.conn.Exec(query,
		game.ID(),
		game.Home(),
		game.Away(),
		game.List().ListID(),
		game.ExternalID(),
		game.Venue(),
		game.Date(),
	)

	return err
}
