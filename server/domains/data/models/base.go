package models

import "github.com/gofrs/uuid"

type BaseModel struct {
	Id uuid.UUID `gorm:"primaryKey;type:uuid.UUID;<-:create"`
	Created
	Updated
}

// Builder Object for BaseModel
type BaseModelBuilder struct {
	id uuid.UUID
	CreatedBuilder
	UpdatedBuilder
}

// Constructor for BaseModelBuilder
func NewBaseModelBuilder() *BaseModelBuilder {
	o := new(BaseModelBuilder)
	return o
}

// Build Method which creates BaseModel
func (b *BaseModelBuilder) Build() *BaseModel {
	o := new(BaseModel)
	o.Id = b.id
	o.Created = *b.CreatedBuilder.Build()
	o.Updated = *b.UpdatedBuilder.Build()
	return o
}

// Setter method for the field id of type uuid.UUID in the object BaseModelBuilder
func (b *BaseModelBuilder) SetId(id uuid.UUID) {
	b.id = id
}
