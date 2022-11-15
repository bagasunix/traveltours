package models

type Country struct {
	Id   int16  `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:200"`
	Coordinate
}

// Builder Object for Country
type CountryBuilder struct {
	id   int16
	name string
	CoordinateBuilder
}

// Constructor for CountryBuilder
func NewCountryBuilder() *CountryBuilder {
	o := new(CountryBuilder)
	return o
}

// Build Method which creates Country
func (b *CountryBuilder) Build() *Country {
	o := new(Country)
	o.Id = b.id
	o.Name = b.name
	o.Coordinate = *b.CoordinateBuilder.Build()
	return o
}

// Setter method for the field id of type int16 in the object CountryBuilder
func (c *CountryBuilder) SetId(id int16) {
	c.id = id
}

// Setter method for the field name of type string in the object CountryBuilder
func (c *CountryBuilder) SetName(name string) {
	c.name = name
}
