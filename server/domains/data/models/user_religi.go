package models

type UserReligi struct {
	BaseModel
	Name string
}

// Builder Object for UserReligi
type UserReligiBuilder struct {
	BaseModelBuilder
	name string
}

// Constructor for UserReligiBuilder
func NewUserReligiBuilder() *UserReligiBuilder {
	o := new(UserReligiBuilder)
	return o
}

// Build Method which creates UserReligi
func (b *UserReligiBuilder) Build() *UserReligi {
	o := new(UserReligi)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object UserReligiBuilder
func (u *UserReligiBuilder) SetName(name string) {
	u.name = name
}
