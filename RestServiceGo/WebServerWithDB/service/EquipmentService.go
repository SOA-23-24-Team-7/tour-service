package service

import (
	"database-example/model"
	"database-example/repo"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) Create(equipment *model.Equipment) error{
	err := service.EquipmentRepo.Create(equipment)
	if(err != nil){
		return err
	}
	return nil
}

func (service *EquipmentService) Find(id int64) (*model.Equipment, error){
	equipment,err := service.EquipmentRepo.Find(id)
	if err != nil{
		return nil, err
	}

	return &equipment, nil
}

func (service *EquipmentService) FindAll()([]model.Equipment,error){
	equipment,err := service.EquipmentRepo.FindAll()
	if err != nil{
		return nil,err
	}
	return equipment,nil
}