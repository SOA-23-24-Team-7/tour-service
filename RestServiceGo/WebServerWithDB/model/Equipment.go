package model

import "errors"

type Equipment struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	//Tours
	Tours 		[]Tour `gorm:"many2many:tour_equipments;"`
}

func (equipment *Equipment) Validate() error {
	if equipment.Name == "" || equipment.Name == " " {
		return errors.New("name can not be empty")
	}
	if equipment.Description == "" || equipment.Description == " "{
		return errors.New("descritprion can not be empty")
	}
	return nil
}