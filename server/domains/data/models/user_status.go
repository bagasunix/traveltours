package models

import "github.com/gofrs/uuid"

type UserStatus struct {
	Id     uuid.UUID `gorm:"primaryKey;type:uuid;<-create"`
	UserId uuid.UUID `gorm:"not null"`
	Name   string    `gorm:"size:20;index:idx_user_status_name,options:CONCURRENTLY;<-create"`
	Created
}

// UserStatusBuilder Builder Object for UserStatus
type UserStatusBuilder struct {
	id     uuid.UUID
	userId uuid.UUID
	name   string
	CreatedBuilder
}

// NewUserStatusBuilder Constructor for UserStatusBuilder
func NewUserStatusBuilder() *UserStatusBuilder {
	o := new(UserStatusBuilder)
	return o
}

// Build Method which creates UserStatus
func (u *UserStatusBuilder) Build() *UserStatus {
	o := new(UserStatus)
	o.Id = u.id
	o.UserId = u.userId
	o.Name = u.name
	o.Created = *u.CreatedBuilder.Build()
	return o
}

// SetId Setter method for the field id of type uuid.UUID in the object UserStatusBuilder
func (u *UserStatusBuilder) SetId(id uuid.UUID) *UserStatusBuilder {
	u.id = id
	return u
}

// SetName Setter method for the field name of type string in the object UserStatusBuilder
func (u *UserStatusBuilder) SetName(name string) *UserStatusBuilder {
	u.name = name
	return u
}

// SetUserId Setter method for the field userId of type uuid.UUID in the object UserStatusBuilder
func (u *UserStatusBuilder) SetUserId(userId uuid.UUID) *UserStatusBuilder {
	u.userId = userId
	return u
}
