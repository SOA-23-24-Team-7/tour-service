package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourPreferenceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourPreferenceRepository) Create(tourPreference *model.TourPreference) error {
	dbResult := repo.DatabaseConnection.Create(tourPreference)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourPreferenceRepository) Find(id int64) (model.TourPreference, error) {
	tourPreference := model.TourPreference{}
	dbResult := repo.DatabaseConnection.Where("user_id = ?", id).First(&tourPreference)
	if dbResult.Error != nil {
		return tourPreference, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return tourPreference, nil
}
