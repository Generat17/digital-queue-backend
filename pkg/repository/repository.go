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
	GetResponsibilityByEmployeeId(employeeId int) ([]string, error)
	GetResponsibilityByWorkstationId(workstationId int) ([]string, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
}

type Repository struct {
	Employee
	Responsibility
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		Employee:       NewEmployeePostgres(db),
		Responsibility: NewResponsibilityPostgres(db),
	}
}
