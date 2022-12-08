package entities

import "time"

type UserDetails struct {
	Entity
	IdentityNo    string    `json:"identity_no"`
	FullName      string    `json:"fullname"`
	Sex           string    `json:"sex"`
	Religi        string    `json:"religi"`
	DOB           time.Time `json:"dob"`
	POB           string    `json:"pob"`
	WhatsApp      string    `json:"whatsapp"`
	Address       string    `json:"address"`
	StreetAddress string    `json:"streetaddress"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	PostalCode    string    `json:"postalcode"`
}

// Builder Object for UserDetails
type UserDetailsBuilder struct {
	EntityBuilder
	identityNo    string
	fullName      string
	sex           string
	religi        string
	dOB           time.Time
	pOB           string
	whatsApp      string
	address       string
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
	o.Entity = *b.EntityBuilder.Build()
	o.IdentityNo = b.identityNo
	o.FullName = b.fullName
	o.Sex = b.sex
	o.Religi = b.religi
	o.DOB = b.dOB
	o.POB = b.pOB
	o.WhatsApp = b.whatsApp
	o.Address = b.address
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
func (u *UserDetailsBuilder) SetSex(sex string) {
	u.sex = sex
}

// Setter method for the field religi of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetReligi(religi string) {
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

// Setter method for the field address of type string in the object UserDetailsBuilder
func (u *UserDetailsBuilder) SetAddress(address string) {
	u.address = address
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
