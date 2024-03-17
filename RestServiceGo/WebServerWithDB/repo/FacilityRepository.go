package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *FacilityRepository) Create(facility *model.Facility) error {
	dbResult := repo.DatabaseConnection.Create(facility)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *FacilityRepository) FindAll(id int64) ([]model.Facility, error) {
	var facilities []model.Facility
	dbResult := repo.DatabaseConnection.Where("author_id = ?", id).Find(&facilities)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return facilities, nil
}
