package service

import (
	"server/pkg/repository"
	"server/types"
)

type Authorization interface {
	CreateEmployee(employee types.Employee) (int, error)
	GenerateToken(login, password string) (string, error)
	GenerateTokenWorkstation(login, password string, workstation int) (string, error)
	ParseToken(token string) (int, error)
	ParseTokenWorkstation(token string) (types.ParseTokenWorkstationResponse, error)
}

type Employee interface {
	GetEmployeeList() ([]types.Employee, error)
}

type Queue interface {
	GetQueueList() ([]types.QueueItem, error)
	AddQueueItem(service string) (int, error)
	GetNewClient(employeeId, workstationId int) (types.GetNewClientResponse, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
}

type Service struct {
	Employee
	Queue
	Responsibility
	Authorization
}

func NewService(repos *repository.Repository, queue *[]types.QueueItem) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		Employee:       NewEmployeeService(repos.Employee),
		Queue:          NewQueueService(repos.Queue, queue),
		Responsibility: NewResponsibilityService(repos.Responsibility),
	}
}
