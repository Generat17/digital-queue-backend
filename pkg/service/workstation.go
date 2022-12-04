package service

import (
	"database/sql"
	"github.com/sirupsen/logrus"
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

	listWorkstationResponsibility, err := s.repo.GetWorkstationResponsibilityList()
	if err != nil {
		return nil, err
	}

	listWorkstation, err := s.repo.GetWorkstationList()
	if err != nil {
		return nil, err
	}

	logrus.Print("Длина массива", len(listWorkstation))
	logrus.Print("Длина друга", len(listWorkstationResponsibility))
	logrus.Print("Длина друга", listWorkstationResponsibility)
	for i := 0; i < len(listWorkstation); i++ {

		for j := 0; j < len(listWorkstationResponsibility); j++ {
			if listWorkstationResponsibility[j].WorkstationId == listWorkstation[i].WorkstationId {
				listWorkstation[i].ResponsibilityList = append(listWorkstation[i].ResponsibilityList, listWorkstationResponsibility[j].ResponsibilityId)
				listWorkstation[i].ResponsibilityListName = append(listWorkstation[i].ResponsibilityListName, listWorkstationResponsibility[j].ResponsibilityName)
			}
		}
	}
	logrus.Print(listWorkstation)
	return listWorkstation, nil
}

// GetWorkstation получает данные о рабочем месте по его ID
func (s *WorkstationService) GetWorkstation(workstationId int) (types.Workstation, error) {
	return s.repo.GetWorkstation(workstationId)
}

// UpdateWorkstation обновляет рабочее место
func (s *WorkstationService) UpdateWorkstation(workstationId int, workstationName string) (sql.Result, error) {
	return s.repo.UpdateWorkstation(workstationId, workstationName)
}

// RemoveWorkstation удаляет рабочее место
func (s *WorkstationService) RemoveWorkstation(workstationId int) (sql.Result, error) {
	return s.repo.RemoveWorkstation(workstationId)
}

// AddWorkstation создает рабочее место
func (s *WorkstationService) AddWorkstation(workstationName string) (sql.Result, error) {
	return s.repo.AddWorkstation(workstationName)
}
