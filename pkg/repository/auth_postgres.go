package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"server/types"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateEmployee(employee types.Employee) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password, first_name, second_name, position) values ($1, $2, $3, $4, $5) RETURNING employee_id", employeeTable)

	row := r.db.QueryRow(query, employee.Login, employee.Password, employee.FirstName, employee.SecondName, employee.Position)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetEmployee(username, password string) (types.Employee, error) {
	var employee types.Employee
	query := fmt.Sprintf("SELECT employee_id FROM %s WHERE login=$1 AND password=$2", employeeTable)
	err := r.db.Get(&employee, query, username, password)

	return employee, err
}
