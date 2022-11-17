package requests

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

type EntityId struct {
	Id any `json:"id"`
}

func (v *EntityId) Validate() error {
	if validation.IsEmpty(v.Id) {
		return errors.ErrInvalidAttributes("id")
	}
	return nil
}

func (s *EntityId) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

// EntityIdBuilder Builder Object for EntityId
type EntityIdBuilder struct {
	id any
}

// NewEntityIdBuilder Constructor for EntityIdBuilder
func NewEntityIdBuilder() *EntityIdBuilder {
	o := new(EntityIdBuilder)
	return o
}

// Build Method which creates EntityId
func (b *EntityIdBuilder) Build() *EntityId {
	o := new(EntityId)
	o.Id = b.id
	return o
}

// SetId Setter method for the field id of type uuid.UUID in the object EntityIdBuilder
func (e *EntityIdBuilder) SetId(id uuid.UUID) *EntityIdBuilder {
	e.id = id
	return e
}
