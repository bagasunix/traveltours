package entities

type UserStatus struct {
	Entity
	Name string `json:"name"`
}

// Builder Object for UserStatus
type UserStatusBuilder struct {
	EntityBuilder
	name string
}

// Constructor for UserStatusBuilder
func NewUserStatusBuilder() *UserStatusBuilder {
	o := new(UserStatusBuilder)
	return o
}

// Build Method which creates UserStatus
func (b *UserStatusBuilder) Build() *UserStatus {
	o := new(UserStatus)
	o.Entity = *b.EntityBuilder.Build()
	o.Name = b.name
	return o
}

// Setter method for the field name of type string in the object UserStatusBuilder
func (u *UserStatusBuilder) SetName(name string) {
	u.name = name
}
