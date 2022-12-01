package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QueuePostgres struct {
	db *sqlx.DB
}

func NewQueuePostgres(db *sqlx.DB) *QueuePostgres {
	return &QueuePostgres{db: db}
}

// GetResponsibilityByEmployeeId получает список обязанностей сотрудника по его ID
func (r *QueuePostgres) GetResponsibilityByEmployeeId(employeeId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM position_responsibility LEFT JOIN responsibility ON position_responsibility.responsibility_id = responsibility.responsibility_id WHERE position_responsibility.position_id = (SELECT position FROM employee WHERE employee_id = $1)")
	err := r.db.Select(&responsibilityName, query, employeeId)

	//select responsibility_name from position_responsibility left join responsibility
	//on position_responsibility.responsibility_id = responsibility.responsibility_id
	//where position_responsibility.position_id = 1;
	return responsibilityName, err
}

// GetResponsibilityByWorkstationId получает функции (обязанности) рабочего места по его ID
func (r *QueuePostgres) GetResponsibilityByWorkstationId(workstationId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM workstation_responsibility LEFT JOIN responsibility ON workstation_responsibility.responsibility_id = responsibility.responsibility_id WHERE workstation_responsibility.workstation_id = $1")
	err := r.db.Select(&responsibilityName, query, workstationId)

	return responsibilityName, err
}

// SetStatusEmployee обновляет статус сотрудника в БД
func (r *QueuePostgres) SetStatusEmployee(statusCode int, employeeId int) (sql.Result, error) {
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE employee_id=$2", employeeTable)
	res, err := r.db.Exec(query, statusCode, employeeId)

	return res, err
}
