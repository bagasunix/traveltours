package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Created struct {
	CreatedBy uuid.UUID `gorm:"type:uuid;index;<-:create"`
	CreatedAt time.Time `gorm:"autoCreateTime;index;sort:desc;<-:create"`
}

// Builder Object for Created
type CreatedBuilder struct {
	createdBy uuid.UUID
	createdAt time.Time
}

// Constructor for CreatedBuilder
func NewCreatedBuilder() *CreatedBuilder {
	o := new(CreatedBuilder)
	return o
}

// Build Method which creates Created
func (b *CreatedBuilder) Build() *Created {
	o := new(Created)
	o.CreatedAt = b.createdAt
	o.CreatedBy = b.createdBy
	return o
}

// Setter method for the field createdBy of type uuid.UUID in the object CreatedBuilder
func (c *CreatedBuilder) SetCreatedBy(createdBy uuid.UUID) *CreatedBuilder {
	c.createdBy = createdBy
	return c
}

// Setter method for the field createdAt of type time.Time in the object CreatedBuilder
func (c *CreatedBuilder) SetCreatedAt(createdAt time.Time) *CreatedBuilder {
	c.createdAt = createdAt
	return c
}
