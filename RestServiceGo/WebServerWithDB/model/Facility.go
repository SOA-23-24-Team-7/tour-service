package model

import (
	"errors"
)

type FacilityCategory int

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
	Longitude   float64
	Latitude    float64
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
