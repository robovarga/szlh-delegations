package repository

import (
	"database/sql"
)

type RefereesRepository struct {
	conn *sql.DB
}

func NewRefereesRepository(conn *sql.DB) *RefereesRepository {
	return &RefereesRepository{conn}
}

func (repo *RefereesRepository) FindBySecureKey() {

}
