package service

import (
	"server/pkg/repository"
	"server/types"
)

type EmployeeService struct {
	repo repository.Employee
}

func NewEmployeeService(repo repository.Employee) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetEmployeeList() ([]types.Employee, error) {

	return s.repo.GetEmployeeList()
}
