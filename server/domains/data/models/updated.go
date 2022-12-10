package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Updated struct {
	UpdatedBy uuid.UUID `gorm:"type:uuid;index;<-:update"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;<-:update"`
}

// UpdatedBuilder Builder Object for Updated
type UpdatedBuilder struct {
	updatedBy uuid.UUID
	updatedAt time.Time
}

// NewUpdatedBuilder Constructor for UpdatedBuilder
func NewUpdatedBuilder() *UpdatedBuilder {
	o := new(UpdatedBuilder)
	return o
}

// Build Method which creates Updated
func (b *UpdatedBuilder) Build() *Updated {
	o := new(Updated)
	o.UpdatedBy = b.updatedBy
	o.UpdatedAt = b.updatedAt
	return o
}

// Setter method for the field updatedBy of type uuid.UUID in the object UpdatedBuilder
func (u *UpdatedBuilder) SetUpdatedBy(updatedBy uuid.UUID) {
	u.updatedBy = updatedBy
}

// Setter method for the field updatedAt of type time.Time in the object UpdatedBuilder
func (u *UpdatedBuilder) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
