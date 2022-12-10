package requests

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type CreateUser struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RoleId   uuid.UUID `json:"role_id"`
}

func (s *CreateUser) Validate() error {
	if validation.IsEmpty(s.Email) {
		return errors.ErrInvalidAttributes("email")
	}
	if validation.IsEmpty(s.Password) {
		return errors.ErrInvalidAttributes("password")
	}
	if validation.IsEmpty(s.RoleId) {
		return errors.ErrInvalidAttributes("role")
	}
	return nil
}

func (s *CreateUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for CreateUser
type CreateUserBuilder struct {
	email    string
	password string
	roleId   uuid.UUID
	statusId uuid.UUID
}

// Constructor for CreateUserBuilder
func NewCreateUserBuilder() *CreateUserBuilder {
	o := new(CreateUserBuilder)
	return o
}

// Build Method which creates CreateUser
func (b *CreateUserBuilder) Build() *CreateUser {
	o := new(CreateUser)
	o.Email = b.email
	o.Password = b.password
	o.RoleId = b.roleId
	return o
}

// Setter method for the field email of type string in the object CreateUserBuilder
func (c *CreateUserBuilder) SetEmail(email string) {
	c.email = email
}

// Setter method for the field password of type string in the object CreateUserBuilder
func (c *CreateUserBuilder) SetPassword(password string) {
	c.password = password
}

// Setter method for the field roleId of type uuid.UUID in the object CreateUserBuilder
func (c *CreateUserBuilder) SetRoleId(roleId uuid.UUID) {
	c.roleId = roleId
}

type DisableAccount struct {
	EntityId
	StatusId int8 `json:"status_id"`
}

func (v *DisableAccount) Validate() error {
	if validation.IsEmpty(v.StatusId) {
		return errors.ErrInvalidAttributes("status")
	}
	return nil
}

func (s *DisableAccount) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for DisableAccount
type DisableAccountBuilder struct {
	EntityIdBuilder
	statusId int8
}

// Constructor for DisableAccountBuilder
func NewDisableAccountBuilder() *DisableAccountBuilder {
	o := new(DisableAccountBuilder)
	return o
}

// Build Method which creates DisableAccount
func (b *DisableAccountBuilder) Build() *DisableAccount {
	o := new(DisableAccount)
	o.EntityId = *b.EntityIdBuilder.Build()
	o.StatusId = b.statusId
	return o
}

// Setter method for the field isActive of type int8 in the object DisableAccountBuilder
func (d *DisableAccountBuilder) SetIsActive(statusId int8) {
	d.statusId = statusId
}
