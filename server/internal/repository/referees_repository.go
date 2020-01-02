package repository

import (
	"database/sql"
	"time"

	"github.com/robovarga/szlh-delegations/internal/entity"
)

type RefereesRepository struct {
	conn *sql.DB
}

func NewRefereesRepository(conn *sql.DB) *RefereesRepository {
	return &RefereesRepository{conn}
}

func (repo *RefereesRepository) GetAll() ([]*entity.Referee, error) {
	query := `SELECT referee_id, name, date_add, date_update FROM referees;`

	rows, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var referees []*entity.Referee

	for rows.Next() {
		var (
			refereeID           int
			name                string
			dateAdd, dateUpdate time.Time
		)

		err = rows.Scan(&refereeID, &name, &dateAdd, &dateUpdate)
		if err != nil {
			// TODO: handle error and continue with processing rows.
			return nil, err
		}

		referees = append(
			referees,
			entity.NewReferee(
				refereeID,
				name,
				dateAdd,
				dateUpdate,
			),
		)
	}

	return referees, nil
}

func (repo *RefereesRepository) InsertRef(name string) (*entity.Referee, error) {
	query := `INSERT INTO referees (name, date_add, date_update) VALUES (?, ?, ?);`

	currentTime := time.Now()

	res, err := repo.conn.Exec(query, name, currentTime, currentTime)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return entity.NewReferee(int(id), name, currentTime, currentTime), nil
}

func (repo *RefereesRepository) FindByName(name string) (*entity.Referee, error) {
	query := `SELECT referee_id, date_add, date_update FROM referees WHERE name = ?`

	var (
		refID               int
		dateAdd, dateUpdate time.Time
	)

	err := repo.conn.QueryRow(query, name).Scan(&refID, &dateAdd, &dateUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return entity.NewReferee(
		refID,
		name,
		dateAdd,
		dateUpdate,
	), nil
}
