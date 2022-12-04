package service

import (
	"database/sql"
	"server/pkg/repository"
	"server/types"
)

type Authorization interface {
	CreateEmployee(employee types.Employee) (int, error)
	GenerateTokenWorkstation(username, password string, workstation int) (string, error)
	ParseTokenWorkstation(token string) (types.ParseTokenWorkstationResponse, error)
	GenerateRefreshToken() (string, error)
	UpdateTokenWorkstation(employeeId, workstationId int, refreshToken string) (string, error)
	SetSession(refreshToken string, workstationId int, employeeId int) (bool, error)
	LogOut(employeeId int) (bool, error)
	GetEmployee(username, password string) (types.Employee, error)
	GetEmployeeById(employeeId int) (types.Employee, error)
	GetStatusEmployee(employeeId int) (int, error)
}

type Employee interface {
	GetEmployeeList() ([]types.Employee, error)
}

type Queue interface {
	GetQueueList() ([]types.QueueItem, error)
	AddQueueItem(service string) (int, error)
	GetNewClient(employeeId, workstationId int) (types.GetNewClientResponse, error)
	ConfirmClient(numberQueue, employeeId int) (int, error)
	EndClient(employeeId int) (int, error)
	SetEmployeeStatus(statusCode, employeeId int) (bool, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
	UpdateResponsibility(responsibilityId int, responsibilityName string) (sql.Result, error)
	RemoveResponsibility(responsibilityId int) (sql.Result, error)
	AddResponsibility(responsibilityName string) (sql.Result, error)
}

type Workstation interface {
	GetWorkstationList() ([]types.Workstation, error)
	GetWorkstation(workstationId int) (types.Workstation, error)
	UpdateWorkstation(workstationId int, workstationName string) (sql.Result, error)
	RemoveWorkstation(workstationId int) (sql.Result, error)
	AddWorkstation(workstationName string) (sql.Result, error)
}

type Service struct {
	Employee
	Queue
	Responsibility
	Authorization
	Workstation
}

func NewService(repos *repository.Repository, queue *[]types.QueueItem) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		Employee:       NewEmployeeService(repos.Employee),
		Queue:          NewQueueService(repos.Queue, queue),
		Responsibility: NewResponsibilityService(repos.Responsibility),
		Workstation:    NewWorkstationService(repos.Workstation),
	}
}
