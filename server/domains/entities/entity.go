package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type Updated struct {
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy uuid.UUID `json:"updatedBy"`
}

type Entity struct {
	Id any `json:"id"`
	Created
}

// EntityBuilder Builder Object for Entity
type EntityBuilder struct {
	id any
	CreatedBuilder
}

// NewEntityBuilder Constructor for EntityBuilder
func NewEntityBuilder() *EntityBuilder {
	o := new(EntityBuilder)
	return o
}

// Build Method which creates Entity
func (b *EntityBuilder) Build() *Entity {
	o := new(Entity)
	o.Id = b.id
	o.Created = *b.CreatedBuilder.Build()
	return o
}

// SetId Setter method for the field id of type uuid.UUID in the object EntityBuilder
func (e *EntityBuilder) SetId(id any) *EntityBuilder {
	e.id = id
	return e
}

type Created struct {
	CreatedBy any       `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreatedBuilder Builder Object for Created
type CreatedBuilder struct {
	createdAt time.Time
	createdBy any
}

// NewCreatedBuilder Constructor for CreatedBuilder
func NewCreatedBuilder() *CreatedBuilder {
	o := new(CreatedBuilder)
	return o
}

// Build Method which creates Created
func (c *CreatedBuilder) Build() *Created {
	o := new(Created)
	o.CreatedAt = c.createdAt
	o.CreatedBy = c.createdBy
	return o
}

// SetCreatedAt Setter method for the field createdAt of type time.Time in the object CreatedBuilder
func (c *CreatedBuilder) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

// SetCreatedBy Setter method for the field createdBy of type any in the object CreatedBuilder
func (c *CreatedBuilder) SetCreatedBy(createdBy any) {
	c.createdBy = createdBy
}
