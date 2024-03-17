package service

import (
	"database-example/model"
	"database-example/repo"
)

type KeyPointService struct {
	KeyPointRepo *repo.KeyPointRepository
}

func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepo.Create(keyPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *KeyPointService) Find(id int64) (*model.KeyPoint, error) {
	keyPoint, err := service.KeyPointRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return &keyPoint, nil
}

func (service *KeyPointService) FindAll(id int64) ([]model.KeyPoint, error) {
	keyPoints, err := service.KeyPointRepo.FindAll(id)
	if err != nil {
		return nil, err
	}
	return keyPoints, nil
}
