package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Created struct {
	CreatedBy uuid.UUID `gorm:"type:uuid.UUID;index;<-:create"`
	CreatedAt time.Time `gorm:"autoCreateTime;Index;sort:desc;<-:create"`
}

// Builder Object for CreatedAt
type CreatedBuilder struct {
	createdBy uuid.UUID
	createdAt time.Time
}

// Constructor for CreatedAtBuilder
func NewCreatedAtBuilder() *CreatedBuilder {
	o := new(CreatedBuilder)
	return o
}

// Build Method which creates CreatedAt
func (b *CreatedBuilder) Build() *Created {
	o := new(Created)
	o.CreatedBy = b.createdBy
	o.CreatedAt = b.createdAt
	return o
}

// Setter method for the field createdBy of type uuid.UUID in the object CreatedAtBuilder
func (c *CreatedBuilder) SetCreatedBy(createdBy uuid.UUID) {
	c.createdBy = createdBy
}

// Setter method for the field createdAt of type time.Time in the object CreatedAtBuilder
func (c *CreatedBuilder) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}
