package requests

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreatePermission struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

func (s *CreatePermission) Validate() error {
	if validation.IsEmpty(s.Name) {
		return errors.ErrInvalidAttributes("permission name")
	}
	if validation.IsEmpty(s.Method) {
		return errors.ErrInvalidAttributes("permission method")
	}
	if validation.IsEmpty(s.Url) {
		return errors.ErrInvalidAttributes("permission url")
	}
	return nil
}

func (s *CreatePermission) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for CreatePermission
type CreatePermissionBuilder struct {
	name   string
	method string
	url    string
}

// Constructor for CreatePermissionBuilder
func NewCreatePermissionBuilder() *CreatePermissionBuilder {
	o := new(CreatePermissionBuilder)
	return o
}

// Build Method which creates CreatePermission
func (b *CreatePermissionBuilder) Build() *CreatePermission {
	o := new(CreatePermission)
	o.Name = b.name
	o.Method = b.method
	o.Url = b.url
	return o
}

// Setter method for the field name of type string in the object CreatePermissionBuilder
func (c *CreatePermissionBuilder) SetName(name string) {
	c.name = name
}

// Setter method for the field method of type string in the object CreatePermissionBuilder
func (c *CreatePermissionBuilder) SetMethod(method string) {
	c.method = method
}

// Setter method for the field url of type string in the object CreatePermissionBuilder
func (c *CreatePermissionBuilder) SetUrl(url string) {
	c.url = url
}

type UpdatePermission struct {
	EntityId
	Name   string `json:"name"`
	Method string `json:"method"`
	Url    string `json:"url"`
}

func (s *UpdatePermission) Validate() error {
	if validation.IsEmpty(s.Id) {
		return errors.ErrInvalidAttributes("permission id")
	}
	return nil
}

func (s *UpdatePermission) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for UpdatePermission
type UpdatePermissionBuilder struct {
	name   string
	method string
	url    string
}

// Constructor for UpdatePermissionBuilder
func NewUpdatePermissionBuilder() *UpdatePermissionBuilder {
	o := new(UpdatePermissionBuilder)
	return o
}

// Build Method which creates UpdatePermission
func (b *UpdatePermissionBuilder) Build() *UpdatePermission {
	o := new(UpdatePermission)
	o.Name = b.name
	o.Method = b.method
	o.Url = b.url
	return o
}

// Setter method for the field name of type string in the object UpdatePermissionBuilder
func (u *UpdatePermissionBuilder) SetName(name string) {
	u.name = name
}

// Setter method for the field method of type string in the object UpdatePermissionBuilder
func (u *UpdatePermissionBuilder) SetMethod(method string) {
	u.method = method
}

// Setter method for the field url of type []string in the object UpdatePermissionBuilder
func (u *UpdatePermissionBuilder) SetUrl(url string) {
	u.url = url
}

type AssignPermissionsToRole struct {
	RoleId        any   `json:"roleId"`
	PermissionIds []any `json:"permissionIds"`
}

func (a *AssignPermissionsToRole) ToJSON() []byte {
	j, err := json.Marshal(a)
	errors.HandlerReturnedVoid(err)
	return j
}

func (a *AssignPermissionsToRole) Validate() error {
	if validation.IsEmpty(a.RoleId) {
		return errors.ErrInvalidAttributes("role id")
	}
	if validation.IsEmpty(a.PermissionIds) {
		return errors.ErrInvalidAttributes("permission ids")
	}
	return nil
}

// AssignPermissionsToRoleBuilder Builder Object for AssignPermissionsToRole
type AssignPermissionsToRoleBuilder struct {
	roleId        any
	permissionIds []any
}

// NewAssignPermissionsBuilder NewAssignPermissionBuilder Constructor for AssignPermissionsToRoleBuilder
func NewAssignPermissionsBuilder() *AssignPermissionsToRoleBuilder {
	o := new(AssignPermissionsToRoleBuilder)
	return o
}

// Build Method which creates AssignPermissionsToRole
func (b *AssignPermissionsToRoleBuilder) Build() *AssignPermissionsToRole {
	o := new(AssignPermissionsToRole)
	o.RoleId = b.roleId
	o.PermissionIds = b.permissionIds
	return o
}

// RoleId Builder method to set the field roleId in AssignPermissionsToRoleBuilder
func (b *AssignPermissionsToRoleBuilder) RoleId(v any) *AssignPermissionsToRoleBuilder {
	b.roleId = v
	return b
}

// PermissionId Builder method to set the field permissionId in AssignPermissionsToRoleBuilder
func (b *AssignPermissionsToRoleBuilder) PermissionId(v []any) *AssignPermissionsToRoleBuilder {
	b.permissionIds = v
	return b
}
