package repo

import (
	"database-example/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
} 

func(repo *TourRepository) Create(tour *model.Tour) error{
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) Find(id int64) (model.Tour, error){
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&tour,id)
	if(dbResult.Error != nil){
		return tour, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tour, nil
}

func (repo *TourRepository) FindAll(id int64) ([]model.Tour,error){
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Where("author_id = ?", id).Find(&tours)

	if(dbResult.Error != nil){
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tours, nil
}

func (repo *TourRepository) AddEquipment(tourId int64,equipment *model.Equipment) error{

	var retrievedEquipment model.Equipment
	tour,errorr := repo.Find(tourId)
	query := `
    SELECT t.*
    FROM equipment t
    JOIN tour_equipments te ON t.id = te.equipment_id
    WHERE te.equipment_id = ?;
`
	//find tour

	errr := repo.DatabaseConnection.Raw(query, equipment.Id).Scan(&retrievedEquipment).Error
	if errr != nil || errorr != nil {
		fmt.Println(errr)
		return errr
	}

	//da li postoji vec veza?
	fmt.Println(retrievedEquipment)
	if retrievedEquipment.Id != 0{
		return errors.New("equipment already in use")
	}
	err := repo.DatabaseConnection.Model(&tour).Association("Equipment").Append(equipment)
	if err != nil {
		return err
	}
	return nil
}