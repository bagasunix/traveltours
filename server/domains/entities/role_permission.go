package entities

import "github.com/gofrs/uuid"

type RolePermission struct {
	RoleId       uuid.UUID `json:"role_id"`
	PermissionId uuid.UUID `json:"permission_id"`
	Created
	Updated
}

// Builder Object for RolePermission
type RolePermissionBuilder struct {
	roleId       uuid.UUID
	permissionId uuid.UUID
	CreatedBuilder
	UpdatedBuilder
}

// Constructor for RolePermissionBuilder
func NewRolePermissionBuilder() *RolePermissionBuilder {
	o := new(RolePermissionBuilder)
	return o
}

// Build Method which creates RolePermission
func (b *RolePermissionBuilder) Build() *RolePermission {
	o := new(RolePermission)
	o.RoleId = b.roleId
	o.PermissionId = b.permissionId
	o.Created = *b.CreatedBuilder.Build()
	o.Updated = *b.UpdatedBuilder.Build()
	return o
}

// Setter method for the field roleId of type uuid.UUID in the object RolePermissionBuilder
func (r *RolePermissionBuilder) SetRoleId(roleId uuid.UUID) {
	r.roleId = roleId
}

// Setter method for the field permissionId of type uuid.UUID in the object RolePermissionBuilder
func (r *RolePermissionBuilder) SetPermissionId(permissionId uuid.UUID) {
	r.permissionId = permissionId
}
