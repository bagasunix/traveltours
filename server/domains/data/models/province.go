package models

type Province struct {
	Id        int16 `gorm:"primaryKey"`
	CountryId int16
	Country   *Country `gorm:"foreignKey:CountryId"`
	Name      string   `gorm:"uniqueIndex;size:50"`
	Coordinate
	PostalCodes string
}

// Builder Object for Province
type ProvinceBuilder struct {
	id        int16
	countryId int16
	country   *Country
	name      string
	CoordinateBuilder
	postalCodes string
}

// Constructor for ProvinceBuilder
func NewProvinceBuilder() *ProvinceBuilder {
	o := new(ProvinceBuilder)
	return o
}

// Build Method which creates Province
func (b *ProvinceBuilder) Build() *Province {
	o := new(Province)
	o.Id = b.id
	o.Country = b.country
	o.CountryId = b.countryId
	o.Name = b.name
	o.Coordinate = *b.CoordinateBuilder.Build()
	o.PostalCodes = b.postalCodes
	return o
}

// Setter method for the field id of type int16 in the object ProvinceBuilder
func (p *ProvinceBuilder) SetId(id int16) {
	p.id = id
}

// Setter method for the field country of type *Country in the object ProvinceBuilder
func (p *ProvinceBuilder) SetCountry(country *Country) {
	p.country = country
	p.countryId = p.country.Id
}

// Setter method for the field name of type string in the object ProvinceBuilder
func (p *ProvinceBuilder) SetName(name string) {
	p.name = name
}

// Setter method for the field postalCodes of type string in the object ProvinceBuilder
func (p *ProvinceBuilder) SetPostalCodes(postalCodes string) {
	p.postalCodes = postalCodes
}
