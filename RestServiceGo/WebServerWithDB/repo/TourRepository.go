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
	dbResult :=  repo.DatabaseConnection.Create(tour)
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

func (repo *TourRepository) AddEquipment(tourId int64,equipmentId int64) error{

	//var retrievedEquipment model.Equipment
	tour,errr := repo.Find(tourId)
	if errr != nil  {
		fmt.Println(errr)
		return errr
	}
	
	var equipment []model.Equipment
	query := `
	SELECT * FROM equipment e
	WHERE e.id = ?;
`
	dbResult := repo.DatabaseConnection.Raw(query, equipmentId).Scan(&equipment)
	 if dbResult.Error != nil{
		fmt.Println(dbResult.Error)
		return errors.New("error fetching equipment")
	 }

	err := repo.DatabaseConnection.Model(&tour).Association("Equipment").Append(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TourRepository) GetToursEquipment(tourId int64) ([]model.Equipment,error){

	var retrievedEquipment []model.Equipment
	query := `
    SELECT t.*
    FROM equipment t
    JOIN tour_equipments te ON t.id = te.equipment_id
    WHERE te.tour_id = ?;
`
	//find tour

	errr := repo.DatabaseConnection.Raw(query, tourId).Scan(&retrievedEquipment).Error
	if errr != nil  {
		fmt.Println(errr)
		return nil,errr
	}

	return retrievedEquipment,nil
}

func (repo *TourRepository) DeleteEquipment(tourId int64,equipmentId int64) error{

	//var retrievedEquipment model.Equipment
	tour,errr := repo.Find(tourId)
	if errr != nil  {
		fmt.Println(errr)
		return errr
	}
	
	var equipment []model.Equipment
	query := `
	SELECT * FROM equipment e
	WHERE e.id = ?;
`
	dbResult := repo.DatabaseConnection.Raw(query, equipmentId).Scan(&equipment)
	 if dbResult.Error != nil{
		fmt.Println(dbResult.Error)
		return errors.New("error fetching equipment")
	 }

	err := repo.DatabaseConnection.Model(&tour).Association("Equipment").Delete(equipment)
	if err != nil {
		return err
	}
	return nil
}