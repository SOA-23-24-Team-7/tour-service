package model

import (
	"errors"

	"github.com/lib/pq"
)

type TourPreference struct {
	Id              int64
	UserId          int64
	DifficultyLevel int
	WalkingRating   int
	CyclingRating   int
	CarRating       int
	BoatRating      int
	SelectedTags    pq.StringArray `gorm:"type:text[]"`
}

func (tp *TourPreference) Validate() error {
	if tp.DifficultyLevel < 1 || tp.DifficultyLevel > 5 {
		return errors.New("Difficulty level must be in range between 1 and 5!")
	}
	if tp.WalkingRating < 0 || tp.CyclingRating < 0 || tp.CarRating < 0 || tp.BoatRating < 0 || tp.WalkingRating > 3 || tp.CyclingRating > 3 || tp.CarRating > 3 || tp.BoatRating > 3 {
		return errors.New("Rating must be value between 0 and 3!")
	}
	return nil
}
