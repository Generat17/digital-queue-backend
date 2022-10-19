package service

import (
	"server/pkg/repository"
	"server/types"
)

type ResponsibilityService struct {
	repo repository.Responsibility
}

func NewResponsibilityService(repo repository.Responsibility) *ResponsibilityService {
	return &ResponsibilityService{repo: repo}
}

func (s *ResponsibilityService) GetResponsibilityList() ([]types.Responsibility, error) {
	return s.repo.GetResponsibilityList()
}
