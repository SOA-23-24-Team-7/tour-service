package model

import (
	"errors"
)

type FacilityCategory int32

const (
	Restaurant FacilityCategory = iota
	ParkingLot
	Toilet
	Hospital
	Cafe
	Pharmacy
	ExchangeOffice
	BusStop
	Shop
	Other
)

type Facility struct {
	Id          int64
	Name        string
	Description string
	ImagePath   string
	AuthorId    int64
	Category    FacilityCategory
	Longitude   float32
	Latitude    float32
}

func (f Facility) Validate() error {
	if f.Name == "" {
		return errors.New("Invalid Name")
	}
	if f.Category < Restaurant || f.Category > Other {
		return errors.New("Invalid Category")
	}
	return nil
}
