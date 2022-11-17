package models

type UserSex struct {
	BaseModel
	Name string
}

// Builder Object for UserSex
type UserSexBuilder struct {
	BaseModelBuilder
	name string
}

// Constructor for UserSexBuilder
func NewUserSexBuilder() *UserSexBuilder {
	o := new(UserSexBuilder)
	return o
}

// Build Method which creates UserSex
func (b *UserSexBuilder) Build() *UserSex {
	o := new(UserSex)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object UserSexBuilder
func (u *UserSexBuilder) SetName(name string) {
	u.name = name
}
