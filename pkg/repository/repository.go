package repository

import (
	"github.com/jmoiron/sqlx"
	"server/types"
)

type Employee interface {
	GetEmployeeList() ([]types.Employee, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
}

type Repository struct {
	Employee
	Responsibility
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee:       NewEmployeePostgres(db),
		Responsibility: NewResponsibilityPostgres(db),
	}
}
