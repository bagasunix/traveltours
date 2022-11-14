package models

type BaseModel struct {
	Id any `gorm:"primaryKey;type:any;<-:create"`
	Created
	Updated
}

// Builder Object for BaseModel
type BaseModelBuilder struct {
	id any
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

// Setter method for the field id of type any in the object BaseModelBuilder
func (b *BaseModelBuilder) SetId(id any) {
	b.id = id
}
