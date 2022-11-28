package requests

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type CreateRole struct {
	Name          string      `json:"name"`
	Group         string      `json:"group"`
	Desc          string      `json:"desc"`
	PermissionIds []uuid.UUID `json:"permissionIds,omitempty"`
	IsActive      string      `json:"is_active"`
}

func (s *CreateRole) Validate() error {
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("role name")
	}
	if validation.IsEmpty(s.Group) {
		return errors.ErrInvalidAttributes("role group")
	}
	return nil
}

func (s *CreateRole) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// CreateRoleBuilder Builder Object for CreateRole
type CreateRoleBuilder struct {
	name          string
	group         string
	desc          string
	permissionIds []uuid.UUID
	isActive      string
}

// NewCreateRoleBuilder Constructor for CreateRoleBuilder
func NewCreateRoleBuilder() *CreateRoleBuilder {
	o := new(CreateRoleBuilder)
	return o
}

// Build Method which creates CreateRole
func (b *CreateRoleBuilder) Build() *CreateRole {
	o := new(CreateRole)
	o.Name = b.name
	o.Group = b.group
	o.Desc = b.desc
	o.PermissionIds = b.permissionIds
	o.IsActive = b.isActive
	return o
}

// SetName SetName Builder method to set the field name in CreateRoleBuilder
func (b *CreateRoleBuilder) SetName(v string) *CreateRoleBuilder {
	b.name = v
	return b
}

// Setter method for the field group of type string in the object CreateRoleBuilder
func (c *CreateRoleBuilder) SetGroup(group string) {
	c.group = group
}

// Setter method for the field desc of type string in the object CreateRoleBuilder
func (c *CreateRoleBuilder) SetDesc(desc string) {
	c.desc = desc
}

// Setter method for the field isActive of type string in the object CreateRoleBuilder
func (c *CreateRoleBuilder) SetIsActive(isActive string) {
	c.isActive = isActive
}

// Setter method for the field permissionIds of type []uuid.UUID in the object CreateRoleBuilder
func (c *CreateRoleBuilder) SetPermissionIds(permissionIds []uuid.UUID) {
	c.permissionIds = permissionIds
}

type UpdateRole struct {
	Id       any
	Name     string `json:"name"`
	Group    string `json:"group"`
	Desc     string `json:"desc"`
	IsActive string `json:"is_active"`
}

func (s *UpdateRole) Validate() error {
	if validation.IsEmpty(s.Id) {
		return errors.ErrInvalidAttributes("role id")
	}
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("role name")
	}
	return nil
}

func (s *UpdateRole) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// UpdateRoleBuilder Builder Object for UpdateRole
type UpdateRoleBuilder struct {
	id       string
	name     string
	group    string
	desc     string
	isActive string
}

// NewUpdateRoleBuilder Constructor for UpdateRoleBuilder
func NewUpdateRoleBuilder() *UpdateRoleBuilder {
	o := new(UpdateRoleBuilder)
	return o
}

// Build Method which creates UpdateRole
func (b *UpdateRoleBuilder) Build() *UpdateRole {
	o := new(UpdateRole)
	o.Id = b.id
	o.Name = b.name
	o.Group = b.group
	o.Desc = b.desc
	o.IsActive = b.isActive
	return o
}

// SetName SetName Builder method to set the field name in UpdateRoleBuilder
func (b *UpdateRoleBuilder) SetName(v string) *UpdateRoleBuilder {
	b.name = v
	return b
}

// SetGroup Group Builder method to set the field group in UpdateRoleBuilder
func (b *UpdateRoleBuilder) SetGroup(v string) *UpdateRoleBuilder {
	b.group = v
	return b
}

// SetId Setter method for the field id of type string in the object UpdateRoleBuilder
func (u *UpdateRoleBuilder) SetId(id string) {
	u.id = id
}

// Setter method for the field desc of type string in the object UpdateRoleBuilder
func (u *UpdateRoleBuilder) SetDesc(desc string) {
	u.desc = desc
}

// Setter method for the field isActive of type string in the object UpdateRoleBuilder
func (u *UpdateRoleBuilder) SetIsActive(isActive string) {
	u.isActive = isActive
}

type RemovePermissionsFromRole struct {
	RoleId        any   `json:"roleId"`
	PermissionIds []any `json:"permissionId"`
}

func (s *RemovePermissionsFromRole) Validate() error {
	if validation.IsEmpty(s.RoleId) {
		return errors.ErrInvalidAttributes("role id")
	}
	if validation.IsEmpty(s.PermissionIds) {
		return errors.ErrInvalidAttributes("permission ids")
	}
	return nil
}

func (s *RemovePermissionsFromRole) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// RemovePermissionsFromRoleBuilder Builder Object for RemovePermissionsFromRole
type RemovePermissionsFromRoleBuilder struct {
	roleId       any
	permissionId []any
}

// NewRemovePermissionsFromRoleBuilder Constructor for RemovePermissionsFromRoleBuilder
func NewRemovePermissionsFromRoleBuilder() *RemovePermissionsFromRoleBuilder {
	o := new(RemovePermissionsFromRoleBuilder)
	return o
}

// Build Method which creates RemovePermissionsFromRole
func (b *RemovePermissionsFromRoleBuilder) Build() *RemovePermissionsFromRole {
	o := new(RemovePermissionsFromRole)
	o.RoleId = b.roleId
	o.PermissionIds = b.permissionId
	return o
}

// RoleId Builder method to set the field roleId in RemovePermissionsFromRoleBuilder
func (b *RemovePermissionsFromRoleBuilder) RoleId(v any) *RemovePermissionsFromRoleBuilder {
	b.roleId = v
	return b
}

// PermissionId Builder method to set the field permissionId in RemovePermissionsFromRoleBuilder
func (b *RemovePermissionsFromRoleBuilder) PermissionId(v []any) *RemovePermissionsFromRoleBuilder {
	b.permissionId = v
	return b
}
