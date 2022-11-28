package models

type Role struct {
	BaseModel
	Name        string       `gorm:"size:200;uniqueIndex:idx_role_name_unique, sort:asc"`
	Group       string       `gorm:"size:100;index:role_group"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	User        *User        `gorm:"foreignKey:RoleId"`
	Desc        string       `gorm:"size:100"`
	IsActive    string       `gorm:"size:1"`
}

// RoleBuilder Builder Object for Role
type RoleBuilder struct {
	BaseModelBuilder
	name        string
	group       string
	permissions []Permission
	desc        string
	isActive    string
}

// NewRoleBuilder Constructor for RoleBuilder
func NewRoleBuilder() *RoleBuilder {
	o := new(RoleBuilder)
	return o
}

// Build Method which creates Role
func (r *RoleBuilder) Build() *Role {
	o := new(Role)
	o.BaseModel = *r.BaseModelBuilder.Build()
	o.Name = r.name
	o.Group = r.group
	o.Permissions = r.permissions
	o.Desc = r.desc
	o.IsActive = r.isActive
	return o
}

// SetName Setter method for the field name of type string in the object RoleBuilder
func (r *RoleBuilder) SetName(name string) *RoleBuilder {
	r.name = name
	return r
}

// SetGroup Setter method for the field group of type string in the object RoleBuilder
func (r *RoleBuilder) SetGroup(group string) *RoleBuilder {
	r.group = group
	return r
}

// SetPermissions Setter method for the field permissions of type []Permission in the object RoleBuilder
func (r *RoleBuilder) SetPermissions(permissions []Permission) *RoleBuilder {
	r.permissions = permissions
	return r
}
func (r *RoleBuilder) SetDesc(desc string) {
	r.desc = desc
}

// Setter method for the field isActive of type string in the object RoleBuilder
func (r *RoleBuilder) SetIsActive(isActive string) {
	r.isActive = isActive
}
