package service

import (
	"database-example/model"
	"database-example/repo"
	"errors"
)

type TourService struct{
	TourRepo *repo.TourRepository
	EquipmentRepo *repo.EquipmentRepository
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

func (service *TourService) FindAll(id int64)([]model.Tour,error){
	tours,err := service.TourRepo.FindAll(id)
	if err != nil{
		return nil,err
	}
	return tours,nil
}

func (service *TourService) AddEquipment(tourId int64, equipmentId int64) error{
	//fetch equipment
	equipment,err := service.EquipmentRepo.Find(equipmentId)
	if err != nil{
		return errors.New("invalid equipmentId")
	}
	err2 := service.TourRepo.AddEquipment(tourId,&equipment)
	if err2 != nil{
		return err2
	}
	return nil
}