package models

type UserStatus struct {
	Id   int8   `gorm:"primaryKey;<-create"`
	Name string `gorm:"size:20;index:idx_user_status_name,options:CONCURRENTLY;<-create"`
	Created
}

// UserStatusBuilder Builder Object for UserStatus
type UserStatusBuilder struct {
	id   int8
	name string
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
	o.Name = u.name
	o.Created = *u.CreatedBuilder.Build()
	return o
}

// SetId Setter method for the field id of type uuid.UUID in the object UserStatusBuilder
func (u *UserStatusBuilder) SetId(id int8) *UserStatusBuilder {
	u.id = id
	return u
}

// SetName Setter method for the field name of type string in the object UserStatusBuilder
func (u *UserStatusBuilder) SetName(name string) *UserStatusBuilder {
	u.name = name
	return u
}
