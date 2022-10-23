package repository

import (
	"github.com/jmoiron/sqlx"
	"server/types"
)

type Authorization interface {
	CreateEmployee(employee types.Employee) (int, error)
	GetEmployee(login, password string) (types.Employee, error)
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
}

type Repository struct {
	Employee
	Responsibility
	Authorization
	Queue
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		Employee:       NewEmployeePostgres(db),
		Responsibility: NewResponsibilityPostgres(db),
		Queue:          NewQueuePostgres(db),
	}
}
