package service

import (
	"database-example/model"
	"database-example/repo"
)

type TourService struct{
	TourRepo *repo.TourRepository
}

func (service *TourService) Create(tour *model.Tour) error{
	err := service.TourRepo.Create(tour)
	if(err != nil){
		return err
	}
	return nil
}

func (service *TourService) Find(id int64) (*model.Tour, error){
	tour,err := service.TourRepo.Find(id)
	if err != nil{
		return nil, err
	}

	return &tour, nil
}