package service

import (
	"server/types"
	"time"
)

type QueueService struct {
	queue []types.QueueItem
}

func NewQueueService(queue []types.QueueItem) *QueueService {
	return &QueueService{queue: queue}
}

func (s *QueueService) GetQueueList() ([]types.QueueItem, error) {
	return s.queue, nil
}

func (s *QueueService) AddQueueItem(service string) (int, error) {
	// Получение текущего времени
	time := (time.Now()).String()

	// Добавляем новые элемент в конец очереди
	s.queue = append(s.queue, types.QueueItem{Id: len(s.queue), Time: time, Service: service})

	return len(s.queue), nil
}
