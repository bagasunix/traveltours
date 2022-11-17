package models

import "github.com/gofrs/uuid"

type User struct {
	BaseModel
	Email      string `gorm:"size:100;uniqueIndex:idx_account_unique"`
	Password   string
	RoleId     uuid.UUID
	Role       *Role        `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Permission []Permission `gorm:"many2many:role_permissions;"`
	IsActive   int8         `gorm:"size:1"`
}

// Builder Object for User
type UserBuilder struct {
	BaseModelBuilder
	email      string
	password   string
	roleId     uuid.UUID
	role       *Role
	permission []Permission
	isActive   int8
}

// Constructor for UserBuilder
func NewUserBuilder() *UserBuilder {
	o := new(UserBuilder)
	return o
}

// Build Method which creates User
func (b *UserBuilder) Build() *User {
	o := new(User)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Email = b.email
	o.Password = b.password
	o.RoleId = b.roleId
	o.Role = b.role
	o.Permission = b.permission
	o.IsActive = b.isActive
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

// Setter method for the field permission of type []Permission in the object UserBuilder
func (u *UserBuilder) SetPermission(permission []Permission) {
	u.permission = permission
}

// Setter method for the field isActive of type int8 in the object UserBuilder
func (u *UserBuilder) SetIsActive(isActive int8) {
	u.isActive = isActive
}
