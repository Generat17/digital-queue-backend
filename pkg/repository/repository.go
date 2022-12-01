package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/types"
)

type Authorization interface {
	CreateEmployee(employee types.Employee) (int, error)
	GetEmployeeId(username, password string) (types.Employee, error)
	GetEmployee(username, password string) (types.Employee, error)
	GetEmployeeById(employeeId int) (types.Employee, error)
	SetSession(refreshToken string, expiresAt int64, workstationId int, employeeId int) (sql.Result, error)
	GetSession(employeeId int) (types.SessionInfo, error)
	ClearSession(employeeId int) (sql.Result, error)
	GetStatusEmployee(employeeId int) (int, error)
}

type Employee interface {
	GetEmployeeList() ([]types.Employee, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
}

type Queue interface {
	GetResponsibilityByEmployeeId(employeeId int) ([]string, error)
	GetResponsibilityByWorkstationId(workstationId int) ([]string, error)
	SetStatusEmployee(statusCode int, employeeId int) (sql.Result, error)
}

type Workstation interface {
	GetWorkstationList() ([]types.Workstation, error)
	GetWorkstation(workstationId int) (types.Workstation, error)
}

type Repository struct {
	Employee
	Responsibility
	Authorization
	Queue
	Workstation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		Employee:       NewEmployeePostgres(db),
		Responsibility: NewResponsibilityPostgres(db),
		Queue:          NewQueuePostgres(db),
		Workstation:    NewWorkstationPostgres(db),
	}
}
