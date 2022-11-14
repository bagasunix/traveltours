package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Updated struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime;<-:update"`
	UpdatedBy uuid.UUID `gorm:"type:uuid.UUID;index;<-:update"`
}

// Builder Object for Updated
type UpdatedBuilder struct {
	updatedAt time.Time
	updatedBy uuid.UUID
}

// Constructor for UpdatedBuilder
func NewUpdatedBuilder() *UpdatedBuilder {
	o := new(UpdatedBuilder)
	return o
}

// Build Method which creates Updated
func (b *UpdatedBuilder) Build() *Updated {
	o := new(Updated)
	o.UpdatedAt = b.updatedAt
	o.UpdatedBy = b.updatedBy
	return o
}

// Setter method for the field updatedAt of type time.Time in the object UpdatedBuilder
func (u *UpdatedBuilder) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

// Setter method for the field updatedBy of type uuid.UUID in the object UpdatedBuilder
func (u *UpdatedBuilder) SetUpdatedBy(updatedBy uuid.UUID) {
	u.updatedBy = updatedBy
}
