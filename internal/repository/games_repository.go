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
	if err != nil {
		return err
	}

	return repo.insertGameRef(game, time.Now())
}

func (repo *GamesRepository) Insert(game *entity.Game) error {
	query := `INSERT INTO games (game_uuid, home_team, away_team, external_id, list_id, venue, game_date, date_add, date_update)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	currentTime := time.Now()

	res, err := repo.conn.Exec(query,
		game.UUID().Bytes(),
		game.Home(),
		game.Away(),
		game.ExternalID(),
		game.List().ListID(),
		game.Venue(),
		game.Date(),
		currentTime,
		currentTime,
	)
	if err != nil {
		return err
	}

	gameID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	game.SetID(int(gameID))

	return repo.insertGameRef(game, currentTime)
}

func (repo *GamesRepository) CheckGame(externalID, listID int) (int, uuid.UUID, error) {
	query := `SELECT game_id, game_uuid FROM games WHERE external_id = ? AND list_id = ?;`

	var (
		gameID   int
		gameUUID uuid.UUID
	)

	row := repo.conn.QueryRow(query, externalID, listID)

	switch err := row.Scan(&gameID, &gameUUID); err {
	case sql.ErrNoRows:
		return 0, uuid.UUID{}, nil
	case nil:
		return gameID, gameUUID, nil
	default:
		return 0, uuid.UUID{}, nil
	}
}

func (repo *GamesRepository) FindByRefId(refId int) ([]*entity.Game, error) {
	query := `SELECT g.game_id, g.game_uuid, g.external_id, g.home_team, g.away_team, g.venue, g.game_date FROM referees r
		LEFT JOIN game_referees gr on r.referee_id = gr.referee_id
		LEFT JOIN games g on gr.game_id = g.game_id
		where r.referee_id = ?`

	rows, err := repo.conn.Query(query, refId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*entity.Game

	for rows.Next() {
		var (
			gameID, externalID int
			gameUUID           uuid.UUID
			home, away, venue  string
			gameDate           time.Time
		)

		err = rows.Scan(&gameID, &gameUUID, &externalID, &home, &away, &venue, &gameDate)
		if err != nil {
			// TODO: handle error and continue with processing rows.
			return nil, err
		}

		game := entity.NewGame(
			gameID,
			gameUUID,
			externalID,
			nil,
			home,
			away,
			venue,
			gameDate,
		)

		games = append(games, game)

	}

	return games, nil
}

func (repo *GamesRepository) FindByListID(listID int) ([]*entity.Game, error) {
	list, err := repo.listRepo.FindByID(listID)
	if err != nil {
		return nil, err
	}

	query := `SELECT game_id, game_uuid, external_id, home_team, away_team, venue, game_date FROM games WHERE list_id = ?;`

	// Execute the query
	rows, err := repo.conn.Query(query, list.ListID())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*entity.Game

	for rows.Next() {
		var (
			gameID, externalID int
			gameUUID           uuid.UUID
			home, away, venue  string
			gameDate           time.Time
		)

		err = rows.Scan(&gameID, &gameUUID, &externalID, &home, &away, &venue, &gameDate)
		if err != nil {
			// TODO: handle error and continue with processing rows.
			return nil, err
		}

		game := entity.NewGame(
			gameID,
			gameUUID,
			externalID,
			list,
			home,
			away,
			venue,
			gameDate,
		)

		queryRefs := `SELECT r.referee_id, r.name, r.date_add, r.date_update FROM game_referees gr
    		LEFT JOIN referees r on gr.referee_id = r.referee_id
    		WHERE gr.game_id = ?;`

		refsRows, err := repo.conn.Query(queryRefs, game.ID())
		if err != nil {
			return nil, err
		}
		defer refsRows.Close()

		for refsRows.Next() {
			var (
				refID                     int
				refName                   string
				refDateAdd, refDateUpdate time.Time
			)

			err = refsRows.Scan(&refID, &refName, &refDateAdd, &refDateUpdate)
			if err != nil {
				// TODO: handle error and continue with processing rows.
				return nil, err
			}

			game.AddReferee(entity.NewReferee(
				refID,
				refName,
				refDateAdd,
				refDateUpdate,
			))

		}

		games = append(games, game)
	}

	return games, nil
}

func (repo *GamesRepository) insertGameRef(game *entity.Game, currentTime time.Time) error {
	queryRemoveAll := `DELETE FROM game_referees WHERE game_id = ?`

	_, err := repo.conn.Exec(queryRemoveAll, game.ID())
	if err != nil {
		return err
	}

	for _, referee := range game.Referees() {
		queryRefRelation := `INSERT INTO game_referees (game_id, referee_id, date_add, date_update) VALUES(?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE date_update = now()`

		_, err := repo.conn.Exec(queryRefRelation, game.ID(), referee.ID(), currentTime, currentTime)
		if err != nil {
			// TODO: inject logger and log error, but continue with inserting.
			return err
		}
	}

	return nil
}
