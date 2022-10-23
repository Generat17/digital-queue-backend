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
