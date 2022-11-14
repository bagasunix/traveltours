package models

import (
	"time"
)

type Updated struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime;<-:update"`
	UpdatedBy any       `gorm:"type:any;index;<-:update"`
}

// Builder Object for Updated
type UpdatedBuilder struct {
	updatedAt time.Time
	updatedBy any
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

// Setter method for the field updatedBy of type any in the object UpdatedBuilder
func (u *UpdatedBuilder) SetUpdatedBy(updatedBy any) {
	u.updatedBy = updatedBy
}
