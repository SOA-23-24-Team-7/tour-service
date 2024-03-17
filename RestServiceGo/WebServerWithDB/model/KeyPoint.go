package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type KeyPoint struct {
	Id              int64
	TourId          int64
	Name            string
	Description     string
	Longitude       float64
	Latitude        float64
	LocationAddress string
	ImagePath       string
	Order           int
	HaveSecret      bool
	KeyPointSecret  KeyPointSecret `gorm:"type:jsonb;"`
}

type KeyPointSecret struct {
	Images      pq.StringArray `gorm:"type:text[]"`
	Description string
}

func (keyPointSecret KeyPointSecret) Value() (driver.Value, error) {
	return json.Marshal(keyPointSecret)
}

func (keyPointSecret *KeyPointSecret) Scan(value interface{}) error {
	if value == nil {
		*keyPointSecret = KeyPointSecret{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, keyPointSecret)
}

func (kp *KeyPoint) Validate() error {
	if kp.TourId == 0 {
		return errors.New("Invalid TourId")
	}
	if kp.Name == "" {
		return errors.New("Invalid Name")
	}
	if kp.Description == "" {
		return errors.New("Invalid Description")
	}
	if kp.Longitude < -180 || kp.Longitude > 180 {
		return errors.New("Invalid Longitude")
	}
	if kp.Latitude < -90 || kp.Latitude > 90 {
		return errors.New("Invalid Latitude")
	}
	if kp.LocationAddress == "" {
		return errors.New("Invalid LocationAddress")
	}
	if kp.ImagePath == "" {
		return errors.New("Invalid ImagePath")
	}

	return nil
}
