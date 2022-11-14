package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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

func (r *WorkstationPostgres) GetWorkstation(workstationId int) (types.Workstation, error) {
	var workstation types.Workstation
	query := fmt.Sprintf("SELECT workstation_id, workstation_name, employee_id FROM %s WHERE workstation_id=$1", workstationTable)
	err := r.db.Get(&workstation, query, workstationId)

	logrus.Print(workstation)
	return workstation, err
}
