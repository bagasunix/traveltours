package entities

type User struct {
	Entity
	Email    string
	Role     *Role  `json:"role,omitempty"`
	IsActive string `json:"is_active"`
}

// Builder Object for User
type UserBuilder struct {
	EntityBuilder
	email    string
	role     *Role
	isActive string
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.Email = b.email
	o.Role = b.role
	o.IsActive = b.isActive
	return o
}

// Setter method for the field email of type string in the object UserBuilder
func (u *UserBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field role of type *Role in the object UserBuilder
func (u *UserBuilder) SetRole(role *Role) {
	u.role = role
}

// Setter method for the field isActive of type string in the object UserBuilder
func (u *UserBuilder) SetIsActive(isActive string) {
	u.isActive = isActive
}
