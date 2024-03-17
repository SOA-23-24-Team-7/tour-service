package service

import (
	"database-example/model"
	"database-example/repo"
)

type FacilityService struct {
	FacilityRepo *repo.FacilityRepository
}

func (service *FacilityService) Create(facility *model.Facility) error {
	err := service.FacilityRepo.Create(facility)
	if err != nil {
		return err
	}
	return nil
}

func (service *FacilityService) FindAll(id int64) ([]model.Facility, error) {
	facilities, err := service.FacilityRepo.FindAll(id)
	if err != nil {
		return nil, err
	}
	return facilities, nil
}
