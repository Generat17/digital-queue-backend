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

// GetQueueList возвращает список клиентов(элементов) в очереди
func (s *QueueService) GetQueueList() ([]types.QueueItem, error) {
	return s.queue, nil
}

// AddQueueItem добавляет нового клиента (элемент) в конец очереди
func (s *QueueService) AddQueueItem(service string) (int, error) {
	// Получение текущего времени
	time := (time.Now()).String()

	// Добавляем новые элемент в конец очереди
	s.queue = append(s.queue, types.QueueItem{Id: len(s.queue), Time: time, Service: service, Workstation: -1, Status: 1})

	return len(s.queue), nil
}

// ConfirmClient подтверждает, что клиент подошел к рабочему месту сотрудника
func (s *QueueService) ConfirmClient(numberQueue, employeeId int) (int, error) {

	copy(s.queue[numberQueue:], s.queue[numberQueue+1:]) // удаляем элемент из очереди
	s.queue[len(s.queue)-1] = types.QueueItem{}          // удаляем элемент из очереди
	s.queue = s.queue[:len(s.queue)-1]                   // удаляем элемент из очереди

	s.repo.SetStatusEmployee(3, employeeId)

	return 3, nil
}

// EndClient завершает обслуживание клиента
func (s *QueueService) EndClient(employeeId int) (int, error) {

	s.repo.SetStatusEmployee(1, employeeId)

	return 1, nil
}

// SetEmployeeStatus устанавливает статус для сотрудника
func (s *QueueService) SetEmployeeStatus(statusCode, employeeId int) (bool, error) {
	s.repo.SetStatusEmployee(statusCode, employeeId)

	return true, nil
}

// GetNewClient Выбирает доступного клиента из очереди и возращает информацию о нем
func (s *QueueService) GetNewClient(employeeId, workstationId int) (types.GetNewClientResponse, error) {
	// получаем список обязанностей сотрудника
	responsibilityEmployeeList, err := s.repo.GetResponsibilityByEmployeeId(employeeId)
	if err != nil {
		return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: "", EmployeeStatus: 1, NumberQueue: 0}, err
	}

	// получаем список обязанностей сотрудника
	responsibilityWorkstationList, err := s.repo.GetResponsibilityByWorkstationId(workstationId)
	if err != nil {
		return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: "", EmployeeStatus: 1, NumberQueue: 0}, err
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
				s.repo.SetStatusEmployee(2, employeeId)
				s.queue[i].Workstation = workstationId // указываем workstation для клиента
				return types.GetNewClientResponse{NumberTicket: s.queue[i].Id, ServiceTicket: s.queue[i].Service, EmployeeStatus: 2, NumberQueue: i}, nil
			}
		}
	}

	return types.GetNewClientResponse{NumberTicket: -1, ServiceTicket: "Нет доступного клиента", EmployeeStatus: 1, NumberQueue: 0}, nil
}
