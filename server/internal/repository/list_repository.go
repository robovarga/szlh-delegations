package repository

import (
	"database/sql"
	"time"

	"github.com/robovarga/szlh-delegations/internal/entity"
)

type ListRepository struct {
	conn *sql.DB
}

func NewListRepository(conn *sql.DB) *ListRepository {
	return &ListRepository{conn}
}

func (repo *ListRepository) FindByID(listID int) (*entity.List, error) {
	query := `SELECT name, url FROM delegation_list WHERE list_id = ?;`

	var (
		name    string
		listURL string
	)

	err := repo.conn.QueryRow(query, listID).Scan(&name, &listURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return entity.NewList(listID, name, listURL), nil
}

func (repo *ListRepository) GetLists() ([]*entity.List, error) {
	query := `SELECT list_id, name, url, date_add, date_update FROM delegation_list ORDER BY date_add DESC;`

	// Execute the query
	rows, err := repo.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []*entity.List

	for rows.Next() {
		var (
			listID              int
			name, listURL       string
			dateAdd, dateUpdate time.Time
		)

		err = rows.Scan(&listID, &name, &listURL, &dateAdd, &dateUpdate)
		if err != nil {
			// TODO: handle error and continue with processing rows.
			return nil, err
		}

		lists = append(lists, entity.NewList(listID, name, listURL))
	}

	return lists, nil
}
