package repo

import (
	"database-example/model"

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