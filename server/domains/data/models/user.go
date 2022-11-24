package models

import "github.com/gofrs/uuid"

type User struct {
	BaseModel
	UserID   uuid.UUID
	Email    string `gorm:"size:100;uniqueIndex:idx_user_unique"`
	Password string
	RoleId   uuid.UUID `gorm:"not null;"`
	Role     *Role     `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Status   string    `gorm:"size:20;not null;index:, type:hash"`
}

// UserBuilder Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	email    string
	password string
	roleId   uuid.UUID
	role     *Role
	status   string
}

// NewUserBuilder Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (u *UserBuilder) Build() *User {
	o := new(User)
	o.BaseModel = *u.BaseModelBuilder.Build()
	o.Email = u.email
	o.Password = u.password
	o.Role = u.role
	o.RoleId = u.roleId
	o.Status = u.status
	return o
}

// Setter method for the field email of type string in the object UserBuilder
func (u *UserBuilder) SetEmail(email string) {
	u.email = email
}

// Setter method for the field password of type string in the object UserBuilder
func (u *UserBuilder) SetPassword(password string) {
	u.password = password
}

// Setter method for the field roleId of type uuid.UUID in the object UserBuilder
func (u *UserBuilder) SetRoleId(roleId uuid.UUID) {
	u.roleId = roleId
}

// Setter method for the field role of type *Role in the object UserBuilder
func (u *UserBuilder) SetRole(role *Role) {
	u.role = role
}

// Setter method for the field status of type string in the object UserBuilder
func (u *UserBuilder) SetStatus(status string) {
	u.status = status
}
