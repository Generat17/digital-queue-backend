package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"server/types"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) GetEmployeeList() ([]types.Employee, error) {
	var employee []types.Employee
	query := fmt.Sprintf("SELECT * FROM %s", employeeTable)
	err := r.db.Select(&employee, query)

	return employee, err
}

func (r *EmployeePostgres) GetResponsibilityByEmployeeId(employeeId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM position_responsibility LEFT JOIN responsibility ON position_responsibility.responsibility_id = responsibility.responsibility_id WHERE position_responsibility.position_id = (SELECT position FROM employee WHERE employee_id = $1)")
	err := r.db.Select(&responsibilityName, query, employeeId)

	//select responsibility_name from position_responsibility left join responsibility
	//on position_responsibility.responsibility_id = responsibility.responsibility_id
	//where position_responsibility.position_id = 1;
	return responsibilityName, err
}

func (r *EmployeePostgres) GetResponsibilityByWorkstationId(workstationId int) ([]string, error) {

	var responsibilityName []string
	query := fmt.Sprintf("SELECT responsibility_name FROM workstation_responsibility LEFT JOIN responsibility ON workstation_responsibility.responsibility_id = responsibility.responsibility_id WHERE workstation_responsibility.workstation_id = $1")
	err := r.db.Select(&responsibilityName, query, workstationId)

	return responsibilityName, err
}
