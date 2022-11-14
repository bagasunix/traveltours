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

// Setter method for the field name of type string in the object GroupMenuBuilder
func (g *GroupMenuBuilder) SetName(name string) {
	g.name = name
}

// Setter method for the field desc of type string in the object GroupMenuBuilder
func (g *GroupMenuBuilder) SetDesc(desc string) {
	g.desc = desc
}

// Setter method for the field isActive of type int8 in the object GroupMenuBuilder
func (g *GroupMenuBuilder) SetIsActive(isActive int8) {
	g.isActive = isActive
}
