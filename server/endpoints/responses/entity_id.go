package responses

import (
	"encoding/json"

	"github.com/bagasunix/traveltours/pkg/errors"
)

type EntityId struct {
	Id any `json:"id"`
}

func (c *EntityId) ToJSON() []byte {
	j, err := json.Marshal(c)
	errors.HandlerReturnedVoid(err)
	return j
}

// Builder Object for EntityId
type EntityIdBuilder struct {
	id any
}

// Constructor for EntityIdBuilder
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

// Setter method for the field id of type any in the object EntityIdBuilder
func (e *EntityIdBuilder) SetId(id any) *EntityIdBuilder {
	e.id = id
	return e
}
