package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserDetails struct {
	BaseModel
	IdentityNo    string
	FullName      string
	SexId         uuid.UUID
	Sex           *UserSex `gorm:"foreignKey:SexId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	ReligiId      uuid.UUID
	Religi        *UserReligi `gorm:"foreignKey:ReligiId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	DOB           time.Time
	POB           string
	WhatsApp      string `gorm:"size:12"`
	VillageId     *int64
	Village       *Village `gorm:"foreignKey:VillageId"`
	StreetAddress string   `gorm:"size:100"`
	Latitude      float64
	Longitude     float64
	PostalCode    string `gorm:"size:6"`
}

// Builder Object for UserDetails
type UserDetailsBuilder struct {
	BaseModelBuilder
	identityNo    string
	fullName      string
	sexId         uuid.UUID
	sex           *UserSex
	religiId      uuid.UUID
	religi        *UserReligi
	dOB           time.Time
	pOB           string
	whatsApp      string
	villageId     *int64
	village       *Village
	streetAddress string
	latitude      float64
	longitude     float64
	postalCode    string
}

// Constructor for UserDetailsBuilder
func NewUserDetailsBuilder() *UserDetailsBuilder {
	o := new(UserDetailsBuilder)
	return o
}

// Build Method which creates UserDetails
func (b *UserDetailsBuilder) Build() *UserDetails {
	o := new(UserDetails)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.IdentityNo = b.identityNo
	o.FullName = b.fullName
	o.SexId = b.sexId
	o.Sex = b.sex
	o.ReligiId = b.religiId
	o.Religi = b.religi
	o.DOB = b.dOB
	o.POB = b.pOB
	o.WhatsApp = b.whatsApp
	o.VillageId = b.villageId
	o.Village = b.village
	o.StreetAddress = b.streetAddress
	o.Latitude = b.latitude
	o.Longitude = b.longitude
	o.PostalCode = b.postalCode
	return o
}

// Setter method for the field identityNo of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetIdentityNo(identityNo string) {
	u.identityNo = identityNo
}

// Setter method for the field fullName of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetFullName(fullName string) {
	u.fullName = fullName
}

// Setter method for the field sex of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetSex(sex *UserSex) {
	u.sex = sex
}

// Setter method for the field religi of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetReligi(religi *UserReligi) {
	u.religi = religi
}

// Setter method for the field dOB of type time.Time in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetDOB(dOB time.Time) {
	u.dOB = dOB
}

// Setter method for the field pOB of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetPOB(pOB string) {
	u.pOB = pOB
}

// Setter method for the field whatsApp of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetWhatsApp(whatsApp string) {
	u.whatsApp = whatsApp
}

// Setter method for the field villageId of type *int64 in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetVillageId(villageId *int64) {
	u.villageId = villageId
}

// Setter method for the field village of type *Village in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetVillage(village *Village) {
	u.village = village
}

// Setter method for the field streetAddress of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetStreetAddress(streetAddress string) {
	u.streetAddress = streetAddress
}

// Setter method for the field latitude of type float64 in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetLatitude(latitude float64) {
	u.latitude = latitude
}

// Setter method for the field longitude of type float64 in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetLongitude(longitude float64) {
	u.longitude = longitude
}

// Setter method for the field postalCode of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetPostalCode(postalCode string) {
	u.postalCode = postalCode
}

// Setter method for the field sexId of type uuid.UUID in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetSexId(sexId uuid.UUID) {
	u.sexId = sexId
}

// Setter method for the field religiId of type uuid.UUID in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetReligiId(religiId uuid.UUID) {
	u.religiId = religiId
}
