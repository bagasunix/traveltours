package models

type SubDistrict struct {
	Id     int32 `gorm:"primaryKey"`
	CityId int16
	City   *City  `gorm:"foreignKey:CityId"`
	Name   string `gorm:"size:100"`
	Coordinate
	PostalCodes string
}

// Builder Object for SubDistrict
type SubDistrictBuilder struct {
	id     int32
	cityId int16
	city   *City
	name   string
	CoordinateBuilder
	postalCodes string
}

// Constructor for SubDistrictBuilder
func NewSubDistrictBuilder() *SubDistrictBuilder {
	o := new(SubDistrictBuilder)
	return o
}

// Build Method which creates SubDistrict
func (b *SubDistrictBuilder) Build() *SubDistrict {
	o := new(SubDistrict)
	o.Id = b.id
	o.CityId = b.cityId
	o.City = b.city
	o.Name = b.name
	o.Coordinate = *b.CoordinateBuilder.Build()
	o.PostalCodes = b.postalCodes
	return o
}

// Setter method for the field id of type int32 in the object SubDistrictBuilder
func (s *SubDistrictBuilder) SetId(id int32) {
	s.id = id
}

// Setter method for the field city of type *City in the object SubDistrictBuilder
func (s *SubDistrictBuilder) SetCity(city *City) {
	s.city = city
	s.cityId = city.Id
}

// Setter method for the field name of type string in the object SubDistrictBuilder
func (s *SubDistrictBuilder) SetName(name string) {
	s.name = name
}

// Setter method for the field postalCodes of type string in the object SubDistrictBuilder
func (s *SubDistrictBuilder) SetPostalCodes(postalCodes string) {
	s.postalCodes = postalCodes
}
