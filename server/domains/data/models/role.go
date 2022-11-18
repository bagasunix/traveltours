package models

type Role struct {
	BaseModel
	Name     string `gorm:"size:50"`
	Group    string `gorm:"size:100;index:role_group"`
	Desc     string `gorm:"size:100"`
	IsActive int8   `gorm:"size:1"`
}

// Builder Object for Role
type RoleBuilder struct {
	BaseModelBuilder
	name     string
	group    string
	desc     string
	isActive int8
}

// Constructor for RoleBuilder
func NewRoleBuilder() *RoleBuilder {
	o := new(RoleBuilder)
	return o
}

// Build Method which creates Role
func (b *RoleBuilder) Build() *Role {
	o := new(Role)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	o.Group = b.group
	o.Desc = b.desc
	o.IsActive = b.isActive
	return o
}

// Setter method for the field name of type string in the object RoleBuilder
func (r *RoleBuilder) SetName(name string) {
	r.name = name
}

// Setter method for the field typeRole of type string in the object RoleBuilder
func (r *RoleBuilder) SetGroup(group string) {
	r.group = group
}

// Setter method for the field desc of type string in the object RoleBuilder
func (r *RoleBuilder) SetDesc(desc string) {
	r.desc = desc
}

// Setter method for the field isActive of type int8 in the object RoleBuilder
func (r *RoleBuilder) SetIsActive(isActive int8) {
	r.isActive = isActive
}
