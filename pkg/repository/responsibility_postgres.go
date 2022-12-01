package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"server/types"
)

type ResponsibilityPostgres struct {
	db *sqlx.DB
}

func NewResponsibilityPostgres(db *sqlx.DB) *ResponsibilityPostgres {
	return &ResponsibilityPostgres{db: db}
}

// GetResponsibilityList получает список всех обязанностей из БД
func (r *ResponsibilityPostgres) GetResponsibilityList() ([]types.Responsibility, error) {
	var responsibility []types.Responsibility
	query := fmt.Sprintf("SELECT * FROM %s", responsibilityTable)
	err := r.db.Select(&responsibility, query)

	return responsibility, err
}
