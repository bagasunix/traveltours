package models

import "github.com/gofrs/uuid"

type LocationTour struct {
	BaseModel
	TourId    uuid.UUID
	Tour      *Tour `gorm:"foreignKey:TourId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	VillageId *int64
	Village   *Village `gorm:"foreignKey:VillageId"`
	Coordinate
	StreetAddress string `gorm:"size:100"`
	Description   string `gorm:"size:100"`
	Day           int
}

// Builder Object for LocationTour
type LocationTourBuilder struct {
	BaseModelBuilder
	tourId    uuid.UUID
	tour      *Tour
	villageId *int64
	village   *Village
	CoordinateBuilder
	streetAddress string
	description   string
	day           int
}

// Constructor for LocationTourBuilder
func NewLocationTourBuilder() *LocationTourBuilder {
	o := new(LocationTourBuilder)
	return o
}

// Build Method which creates LocationTour
func (b *LocationTourBuilder) Build() *LocationTour {
	o := new(LocationTour)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.TourId = b.tourId
	o.Tour = b.tour
	o.VillageId = b.villageId
	o.Village = b.village
	o.Coordinate = *b.CoordinateBuilder.Build()
	o.StreetAddress = b.streetAddress
	o.Description = b.description
	o.Day = b.day
	return o
}

// Setter method for the field villageId of type *int64 in the object LocationTourBuilder
func (l *LocationTourBuilder) SetVillageId(villageId *int64) {
	l.villageId = villageId
}

// Setter method for the field village of type *Village in the object LocationTourBuilder
func (l *LocationTourBuilder) SetVillage(village *Village) {
	l.village = village
}

// Setter method for the field streetAddress of type string in the object LocationTourBuilder
func (l *LocationTourBuilder) SetStreetAddress(streetAddress string) {
	l.streetAddress = streetAddress
}

// Setter method for the field description of type string in the object LocationTourBuilder
func (l *LocationTourBuilder) SetDescription(description string) {
	l.description = description
}

// Setter method for the field day of type int in the object LocationTourBuilder
func (l *LocationTourBuilder) SetDay(day int) {
	l.day = day
}

// Setter method for the field tourId of type uuid.UUID in the object LocationTourBuilder
func (l *LocationTourBuilder) SetTourId(tourId uuid.UUID) {
	l.tourId = tourId
}

// Setter method for the field tour of type *Tour in the object LocationTourBuilder
func (l *LocationTourBuilder) SetTour(tour *Tour) {
	l.tour = tour
}
