package repository

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/robovarga/szlh-delegations/internal/entity"
)

type GamesRepository struct {
	conn     *sql.DB
	listRepo *ListRepository
}

func NewGamesRepository(
	conn *sql.DB,
	listRepository *ListRepository,
) *GamesRepository {

	return &GamesRepository{conn, listRepository}
}

func (repo *GamesRepository) Update(game *entity.Game) error {
	query := `UPDATE games SET home_team = ?, away_team = ?, venue = ?, game_date = ?, date_update = now()
		WHERE external_id = ? AND list_id = ?;`

	_, err := repo.conn.Exec(query,
		game.Home(),
		game.Away(),
		game.Venue(),
		game.Date(),
		game.ExternalID(),
		game.List().ListID(),
	)

	return err
}

func (repo *GamesRepository) Insert(game *entity.Game) error {
	query := `INSERT INTO games (game_id, home_team, away_team, external_id, list_id, venue, game_date, date_add, date_update)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := repo.conn.Exec(query,
		game.ID(),
		game.Home(),
		game.Away(),
		game.ExternalID(),
		game.List().ListID(),
		game.Venue(),
		game.Date(),
		time.Now(),
		time.Now(),
	)

	return err
}

func (repo *GamesRepository) CheckGame(externalID, listID int) (bool, error) {
	query := `SELECT game_id FROM games WHERE external_id = ? AND list_id = ?;`

	var gameID uuid.UUID

	row := repo.conn.QueryRow(query, externalID, listID)

	switch err := row.Scan(&gameID); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, nil
	}
}

func (repo *GamesRepository) FindByListID(listID int) ([]*entity.Game, error) {
	list, err := repo.listRepo.FindByID(listID)
	if err != nil {
		return nil, err
	}

	query := `SELECT game_id, external_id, home_team, away_team, venue, game_date FROM games WHERE list_id = ?;`

	// Execute the query
	rows, err := repo.conn.Query(query, list.ListID())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*entity.Game

	for rows.Next() {
		var (
			gameID            uuid.UUID
			externalID        int
			home, away, venue string
			gameDate          time.Time
		)

		err = rows.Scan(&gameID, &externalID, &home, &away, &venue, &gameDate)
		if err != nil {
			// TODO: handle error and continue with processing rows.
			return nil, err
		}

		games = append(
			games,
			entity.NewGame(
				gameID,
				externalID,
				list,
				home,
				away,
				venue,
				gameDate,
			),
		)
	}

	return games, nil
}
