package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"server/types"
)

type WorkstationPostgres struct {
	db *sqlx.DB
}

func NewWorkstationPostgres(db *sqlx.DB) *WorkstationPostgres {
	return &WorkstationPostgres{db: db}
}

func (r *WorkstationPostgres) GetWorkstationList() ([]types.Workstation, error) {
	var workstation []types.Workstation
	query := fmt.Sprintf("SELECT * FROM %s", workstationTable)
	err := r.db.Select(&workstation, query)

	return workstation, err
}
