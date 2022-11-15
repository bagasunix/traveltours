package models

type Village struct {
	Id            int64 `gorm:"primaryKey"`
	SubDistrictId int32
	SubDistrict   *SubDistrict `gorm:"foreignKey:SubDistrictId"`
	Name          string       `gorm:"size:100;index"`
	Coordinate
	PostalCodes string  `gorm:"size:500;index"`
	Distance    float64 `gorm:"-"`
}

// VillageBuilder Builder Object for Village
type VillageBuilder struct {
	id            int64
	subDistrictId int32
	subDistrict   *SubDistrict
	name          string
	CoordinateBuilder
	postalCodes string
	distance    float64
}

// NewVillageBuilder Constructor for VillageBuilder
func NewVillageBuilder() *VillageBuilder {
	o := new(VillageBuilder)
	return o
}

// Build Method which creates Village
func (v *VillageBuilder) Build() *Village {
	o := new(Village)

	o.Id = v.id
	o.SubDistrict = v.subDistrict
	o.SubDistrictId = v.subDistrictId
	o.Name = v.name
	o.Coordinate = *v.CoordinateBuilder.Build()
	o.PostalCodes = v.postalCodes
	o.Distance = v.distance
	return o
}

// SetId Setter method for the field id of type int32 in the object VillageBuilder
func (v *VillageBuilder) SetId(id int64) {
	v.id = id
}

// SetSubDistrict Setter method for the field subDistrict of type *SubDistrict in the object VillageBuilder
func (v *VillageBuilder) SetSubDistrict(subDistrict *SubDistrict) {
	v.subDistrict = subDistrict
	v.subDistrictId = v.subDistrict.Id
}

// SetName Setter method for the field name of type string in the object VillageBuilder
func (v *VillageBuilder) SetName(name string) {
	v.name = name
}

// SetPostalCodes Setter method for the field postalCodes of type string in the object VillageBuilder
func (v *VillageBuilder) SetPostalCodes(postalCodes string) {
	v.postalCodes = postalCodes
}

// SetDistance Setter method for the field distance of type float64 in the object VillageBuilder
func (v *VillageBuilder) SetDistance(distance float64) {
	v.distance = distance
}
