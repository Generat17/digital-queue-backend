package service

import (
	"server/pkg/repository"
	"server/types"
	"time"
)

type QueueService struct {
	repo  repository.Queue
	queue []types.QueueItem
}

func NewQueueService(repo repository.Queue, queue *[]types.QueueItem) *QueueService {
	return &QueueService{repo: repo, queue: *queue}
}

func (s *QueueService) GetQueueList() ([]types.QueueItem, error) {
	return s.queue, nil
}

func (s *QueueService) AddQueueItem(service string) (int, error) {
	// Получение текущего времени
	time := (time.Now()).String()

	// Добавляем новые элемент в конец очереди
	s.queue = append(s.queue, types.QueueItem{Id: len(s.queue), Time: time, Service: service, Workstation: -1, Status: 1})

	return len(s.queue), nil
}

func (s *QueueService) GetNewClient(employeeId, workstationId int) (types.GetNewClientResponse, error) {
	// получаем список обязанностей сотрудника
	responsibilityEmployeeList, err := s.repo.GetResponsibilityByEmployeeId(employeeId)
	if err != nil {
		return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: ""}, err
	}

	// получаем список обязанностей сотрудника
	responsibilityWorkstationList, err := s.repo.GetResponsibilityByEmployeeId(workstationId)
	if err != nil {
		return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: ""}, err
	}

	// находим общие обязанности у сотрудника и рабочего места
	var generalResponsibility = []string{}
	for i := 0; i < len(responsibilityEmployeeList); i++ {
		for j := 0; j < len(responsibilityWorkstationList); j++ {
			if responsibilityEmployeeList[i] == responsibilityWorkstationList[j] {
				generalResponsibility = append(generalResponsibility, responsibilityEmployeeList[i])
			}
		}
	}

	// Пробегаемся по очереди и смотрим кого можно принять
	for i := 0; i < len(s.queue); i++ {
		for j := 0; j < len(generalResponsibility); j++ {
			if (s.queue[i].Status == 1) && (s.queue[i].Service == generalResponsibility[j]) {
				s.queue[i].Status = 2                  // изменяем статус клиента
				s.queue[i].Workstation = workstationId // указываем workstation для клиента
				return types.GetNewClientResponse{NumberTicket: s.queue[i].Id, ServiceTicket: s.queue[i].Service}, nil
			}
		}
	}

	return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: "nothing"}, nil
}
