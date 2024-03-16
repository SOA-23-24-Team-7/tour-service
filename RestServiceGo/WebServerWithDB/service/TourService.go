package service

import (
	"database-example/model"
	"database-example/repo"
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

func  (service *TourService) AddEquipment(tourId int64, equipmentId int64) error{
	//fetch equipment
	
	err2 := service.TourRepo.AddEquipment(tourId,equipmentId)
	if err2 != nil{
		return err2
	}
	return nil
}

func (service *TourService) GetEquipment(tourId int64) ([]model.Equipment,error){
	//fetch equipment
	
	equioment,err2 := service.TourRepo.GetToursEquipment(tourId)
	if err2 != nil{
		return nil,err2
	}
	return equioment,nil
}

func  (service *TourService) DeleteEquipment(tourId int64, equipmentId int64) error{

	err2 := service.TourRepo.DeleteEquipment(tourId,equipmentId)
	if err2 != nil{
		return err2
	}
	return nil
}