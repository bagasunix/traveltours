package models

import "github.com/gofrs/uuid"

type User struct {
	BaseModel
	Email    string `gorm:"size:100;uniqueIndex:idx_user_unique"`
	Password string
	RoleId   uuid.UUID   `gorm:"not null;"`
	Role     *Role       `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	StatusId int8        `gorm:"not null;"`
	Status   *UserStatus `gorm:"foreignKey:StatusId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
}

// UserBuilder Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	email    string
	password string
	roleId   uuid.UUID
	role     *Role
	statusId int8
	status   *UserStatus
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
	o.StatusId = u.statusId
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

// Setter method for the field statusId of type uuid.UUID in the object UserBuilder
func (u *UserBuilder) SetStatusId(statusId int8) {
	u.statusId = statusId
}

// Setter method for the field status of type *UserStatus in the object UserBuilder
func (u *UserBuilder) SetStatus(status *UserStatus) {
	u.status = status
}
