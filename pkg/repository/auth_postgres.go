package repository

import (
	"database/sql"
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
	query := fmt.Sprintf("INSERT INTO %s (username, password, first_name, second_name, position) values ($1, $2, $3, $4, $5) RETURNING employee_id", employeeTable)

	row := r.db.QueryRow(query, employee.Username, employee.Password, employee.FirstName, employee.SecondName, employee.Position)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetEmployeeId(username, password string) (types.Employee, error) {
	var employee types.Employee
	query := fmt.Sprintf("SELECT employee_id FROM %s WHERE username=$1 AND password=$2", employeeTable)
	err := r.db.Get(&employee, query, username, password)

	return employee, err
}

func (r *AuthPostgres) GetEmployeeById(employeeId int) (types.Employee, error) {
	var employee types.Employee
	query := fmt.Sprintf("SELECT * FROM %s WHERE employee_id=$1", employeeTable)
	err := r.db.Get(&employee, query, employeeId)

	return employee, err
}

func (r *AuthPostgres) GetEmployee(username, password string) (types.Employee, error) {
	var employee types.Employee
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password=$2", employeeTable)
	err := r.db.Get(&employee, query, username, password)

	return employee, err
}

func (r *AuthPostgres) SetSession(refreshToken string, expiresAt int64, workstationId int, employeeId int) (sql.Result, error) {

	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1, expires_at=$2, workstation_id=$3 WHERE employee_id=$4", employeeTable)
	res, err := r.db.Exec(query, refreshToken, expiresAt, workstationId, employeeId)

	return res, err
}

func (r *AuthPostgres) CheckSession(employeeId int) (types.SessionInfo, error) {

	var sessionInfo types.SessionInfo
	query := fmt.Sprintf("SELECT refresh_token, expires_at, workstation_id FROM %s WHERE employee_id=$1", employeeTable)
	err := r.db.Get(&sessionInfo, query, employeeId)

	return sessionInfo, err
}

func (r *AuthPostgres) ClearSession(employeeId int) (sql.Result, error) {
	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1, expires_at=$2, workstation_id=$3 WHERE employee_id=$4", employeeTable)
	res, err := r.db.Exec(query, "", 0, -1, employeeId)

	return res, err
}
