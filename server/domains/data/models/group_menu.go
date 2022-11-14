package models

type GroupMenu struct {
	BaseModel
	Name     string
	Desc     string
	IsActive int8
}

// Builder Object for GroupMenu
type GroupMenuBuilder struct {
	BaseModelBuilder
	name     string
	desc     string
	isActive int8
}

// Constructor for GroupMenuBuilder
func NewGroupMenuBuilder() *GroupMenuBuilder {
	o := new(GroupMenuBuilder)
	return o
}

// Build Method which creates GroupMenu
func (b *GroupMenuBuilder) Build() *GroupMenu {
	o := new(GroupMenu)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	o.Desc = b.desc
	o.IsActive = b.isActive
	return o
}
