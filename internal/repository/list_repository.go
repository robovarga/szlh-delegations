package repository

import (
	"database/sql"

	"github.com/robovarga/szlh-delegations/internal/entity"
)

type ListRepository struct {
	conn *sql.DB
}

func NewListRepository(conn *sql.DB) *ListRepository {
	return &ListRepository{conn}
}

func (repo *ListRepository) FindByID(listID int) (*entity.List, error) {
	var (
		name    string
		listURL string
	)

	query := `SELECT name, url FROM delegation_list WHERE list_id = ?`

	err := repo.conn.QueryRow(query, listID).Scan(&name, &listURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return entity.NewList(listID, name, listURL), nil
}
