package entities

type User struct {
	Entity
	Email  string `json:"email"`
	Role   string `json:"role,omitempty"`
	Status string `json:"status"`
}

// Builder Object for User
type UserBuilder struct {
	EntityBuilder
	email  string
	role   string
	status string
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.Entity = *b.EntityBuilder.Build()
	o.Email = b.email
	o.Role = b.role
	o.Status = b.status
	return o
}

// Setter method for the field email of type string in the object UserBuilder
func (u *UserBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field role of type string in the object UserBuilder
func (u *UserBuilder) SetRole(role string) {
	u.role = role
}

// Setter method for the field status of type string in the object UserBuilder
func (u *UserBuilder) SetStatus(status string) {
	u.status = status
}
