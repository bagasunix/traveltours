package models

type City struct {
	Id         int16 `gorm:"primaryKey"`
	ProvinceId int16
	Province   *Province `gorm:"foreignKey:ProvinceId"`
	Name       string    `gorm:"size:100"`
	Coordinate
	PostalCodes string
}

// Builder Object for City
type CityBuilder struct {
	id         int16
	provinceId int16
	province   *Province
	name       string
	CoordinateBuilder
	postalCodes string
}

// Constructor for CityBuilder
func NewCityBuilder() *CityBuilder {
	o := new(CityBuilder)
	return o
}

// Build Method which creates City
func (b *CityBuilder) Build() *City {
	o := new(City)
	o.Id = b.id
	o.Province = b.province
	o.ProvinceId = b.provinceId
	o.Name = b.name
	o.Coordinate = *b.CoordinateBuilder.Build()
	o.PostalCodes = b.postalCodes
	return o
}

// Setter method for the field id of type int16 in the object CityBuilder
func (c *CityBuilder) SetId(id int16) {
	c.id = id
}

// Setter method for the field province of type *Province in the object CityBuilder
func (c *CityBuilder) SetProvince(province *Province) {
	c.province = province
	c.provinceId = province.Id
}

// Setter method for the field name of type string in the object CityBuilder
func (c *CityBuilder) SetName(name string) {
	c.name = name
}

// Setter method for the field postalCodes of type string in the object CityBuilder
func (c *CityBuilder) SetPostalCodes(postalCodes string) {
	c.postalCodes = postalCodes
}
