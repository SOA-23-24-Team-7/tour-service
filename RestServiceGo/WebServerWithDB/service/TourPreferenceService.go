package service

import (
	"database-example/model"
	"database-example/repo"
)

type TourPreferenceService struct {
	TourPreferenceRepo *repo.TourPreferenceRepository
}

func (service *TourPreferenceService) Create(tourPreference *model.TourPreference) error {
	err := service.TourPreferenceRepo.Create(tourPreference)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPreferenceService) Find(id int64) (*model.TourPreference, error) {
	tourPreference, err := service.TourPreferenceRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return &tourPreference, nil
}
