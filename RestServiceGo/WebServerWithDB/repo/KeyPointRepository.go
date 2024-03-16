package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
} 

func(repo *KeyPointRepository) Create(keyPoint *model.KeyPoint) error{
	dbResult := repo.DatabaseConnection.Create(keyPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *KeyPointRepository) Find(id int64) (model.KeyPoint, error){
	keyPoint := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&keyPoint,id)
	if(dbResult.Error != nil){
		return keyPoint, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return keyPoint, nil
}

func (repo *KeyPointRepository) FindAll(id int64) ([]model.KeyPoint,error){
	var keyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Where("tour_id = ?", id).Find(&keyPoints)

	if(dbResult.Error != nil){
		return nil, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return keyPoints, nil
}
