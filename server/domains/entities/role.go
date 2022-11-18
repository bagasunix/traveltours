package entities

type Role struct {
	Entity
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	Group       string       `json:"group"`
	IsEnable    bool         `json:"isEnable"`
}

// RoleBuilder Builder Object for Role
type RoleBuilder struct {
	EntityBuilder
	name        string
	permissions []Permission
	group       string
	isEnable    bool
}

// NewRoleBuilder Constructor for RoleBuilder
func NewRoleBuilder() *RoleBuilder {
	o := new(RoleBuilder)
	return o
}

// Build Method which creates Role
func (b *RoleBuilder) Build() *Role {
	o := new(Role)
	o.Entity = *b.EntityBuilder.Build()
	o.Name = b.name
	o.Permissions = b.permissions
	o.Group = b.group
	o.IsEnable = b.isEnable
	return o
}

// SetName Setter method for the field name of type string in the object RoleBuilder
func (b *RoleBuilder) SetName(name string) *RoleBuilder {
	b.name = name
	return b
}

// SetPermissions Setter method for the field permissions of type []Permission in the object RoleBuilder
func (b *RoleBuilder) SetPermissions(permissions []Permission) *RoleBuilder {
	b.permissions = permissions
	return b
}

// SetGroup Setter method for the field group of type string in the object RoleBuilder
func (b *RoleBuilder) SetGroup(group string) *RoleBuilder {
	b.group = group
	return b
}

// SetIsEnable Setter method for the field isEnable of type bool in the object RoleBuilder
func (b *RoleBuilder) SetIsEnable(isEnable bool) *RoleBuilder {
	b.isEnable = isEnable
	return b
}
