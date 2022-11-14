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

func (s *WorkstationService) GetWorkstationList() ([]types.Workstation, error) {
	return s.repo.GetWorkstationList()
}

func (s *WorkstationService) GetWorkstation(workstationId int) (types.Workstation, error) {
	return s.repo.GetWorkstation(workstationId)
}
