package service

import (
	"server/pkg/repository"
	"server/types"
)

type Employee interface {
	GetEmployeeList() ([]types.Employee, error)
}

type Queue interface {
	GetQueueList() ([]types.QueueItem, error)
	AddQueueItem(service string) (int, error)
}

type Responsibility interface {
	GetResponsibilityList() ([]types.Responsibility, error)
}

type Service struct {
	Employee
	Queue
	Responsibility
}

func NewService(repos *repository.Repository, queue *[]types.QueueItem) *Service {
	return &Service{
		Employee:       NewEmployeeService(repos.Employee),
		Queue:          NewQueueService(*queue),
		Responsibility: NewResponsibilityService(repos.Responsibility),
	}
}
