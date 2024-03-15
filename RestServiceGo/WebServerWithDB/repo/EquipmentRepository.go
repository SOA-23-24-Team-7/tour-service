package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func(repo *EquipmentRepository) Create(equipment *model.Equipment) error{
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) Find(id int64) (model.Equipment, error){
	equipment := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&equipment,id)
	if(dbResult.Error != nil){
		return equipment, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return equipment, nil
}

func (repo *EquipmentRepository) FindAll() ([]model.Equipment,error){
	var equipment []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&equipment)

	if(dbResult.Error != nil){
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return equipment, nil
}