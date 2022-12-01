package service

import (
	"server/pkg/repository"
	"server/types"
)

type WorkstationService struct {
	repo repository.Workstation
}

func NewWorkstationService(repo repository.Workstation) *WorkstationService {
	return &WorkstationService{repo: repo}
}

// GetWorkstationList получает список рабочих мест
func (s *WorkstationService) GetWorkstationList() ([]types.Workstation, error) {
	return s.repo.GetWorkstationList()
}

// GetWorkstation получает данные о рабочем месте по его ID
func (s *WorkstationService) GetWorkstation(workstationId int) (types.Workstation, error) {
	return s.repo.GetWorkstation(workstationId)
}
