package models

import (
	"github.com/gofrs/uuid"
)

type RolePermission struct {
	RoleId       uuid.UUID `gorm:"primaryKey"`
	PermissionId uuid.UUID `gorm:"primaryKey"`
	Created
}

// RolePermissionBuilder Builder Object for RolePermission
type RolePermissionBuilder struct {
	roleId       uuid.UUID
	permissionId uuid.UUID
	CreatedBuilder
}

// NewRolePermissionBuilder Constructor for RolePermissionBuilder
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
	return o
}

// SetRoleId Setter method for the field roleId of type uuid.UUID in the object RolePermissionBuilder
func (b *RolePermissionBuilder) SetRoleId(roleId uuid.UUID) *RolePermissionBuilder {
	b.roleId = roleId
	return b
}

// SetPermissionId Setter method for the field permissionId of type uuid.UUID in the object RolePermissionBuilder
func (b *RolePermissionBuilder) SetPermissionId(permissionId uuid.UUID) *RolePermissionBuilder {
	b.permissionId = permissionId
	return b
}
