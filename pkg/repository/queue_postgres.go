package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QueuePostgres struct {
	db *sqlx.DB
}

func NewQueuePostgres(db *sqlx.DB) *QueuePostgres {
	return &QueuePostgres{db: db}
}

func (r *QueuePostgres) GetResponsibilityByEmployeeId(employeeId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM position_responsibility LEFT JOIN responsibility ON position_responsibility.responsibility_id = responsibility.responsibility_id WHERE position_responsibility.position_id = (SELECT position FROM employee WHERE employee_id = $1)")
	err := r.db.Select(&responsibilityName, query, employeeId)

	//select responsibility_name from position_responsibility left join responsibility
	//on position_responsibility.responsibility_id = responsibility.responsibility_id
	//where position_responsibility.position_id = 1;
	return responsibilityName, err
}

func (r *QueuePostgres) GetResponsibilityByWorkstationId(workstationId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM workstation_responsibility LEFT JOIN responsibility ON workstation_responsibility.responsibility_id = responsibility.responsibility_id WHERE workstation_responsibility.workstation_id = $1")
	err := r.db.Select(&responsibilityName, query, workstationId)

	return responsibilityName, err
}
