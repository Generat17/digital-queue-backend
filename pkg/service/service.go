package service

import (
	"server/pkg/repository"
	"server/types"
)

type Authorization interface {
	CreateEmployee(employee types.Employee) (int, error)
	GenerateTokenWorkstation(username, password string, workstation int) (string, error)
	ParseTokenWorkstation(token string) (types.ParseTokenWorkstationResponse, error)
	GenerateRefreshToken() (string, error)
	UpdateTokenWorkstation(employeeId, workstationId int, refreshToken string) (string, error)
	SetSession(refreshToken string, employeeId int) (bool, error)
	GetEmployee(username, password string) (types.Employee, error)
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
