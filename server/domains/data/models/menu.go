package models

import (
	"github.com/gofrs/uuid"
)

type Menu struct {
	BaseModel
	GroupMenuId uuid.UUID
	GroupMenu   *GroupMenu `gorm:"foreignKey:GroupMenuId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Title       string
	Url         string
	Icon        string
	Desc        string
	IsActive    int8
}

// Builder Object for Menu
type MenuBuilder struct {
	BaseModelBuilder
	groupMenuId uuid.UUID
	groupMenu   *GroupMenu
	title       string
	url         string
	icon        string
	desc        string
	isActive    int8
}

// Constructor for MenuBuilder
func NewMenuBuilder() *MenuBuilder {
	o := new(MenuBuilder)
	return o
}

// Build Method which creates Menu
func (b *MenuBuilder) Build() *Menu {
	o := new(Menu)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.GroupMenuId = b.groupMenuId
	o.GroupMenu = b.groupMenu
	o.Title = b.title
	o.Url = b.url
	o.Icon = b.icon
	o.Desc = b.desc
	o.IsActive = b.isActive
	return o
}

// Setter method for the field groupMenuId of type uuid.UUID in the object MenuBuilder
func (m *MenuBuilder) SetGroupMenuId(groupMenuId uuid.UUID) {
	m.groupMenuId = groupMenuId
}

// Setter method for the field groupMenu of type *GroupMenu in the object MenuBuilder
func (m *MenuBuilder) SetGroupMenu(groupMenu *GroupMenu) {
	m.groupMenu = groupMenu
}

// Setter method for the field title of type string in the object MenuBuilder
func (m *MenuBuilder) SetTitle(title string) {
	m.title = title
}

// Setter method for the field url of type string in the object MenuBuilder
func (m *MenuBuilder) SetUrl(url string) {
	m.url = url
}

// Setter method for the field icon of type string in the object MenuBuilder
func (m *MenuBuilder) SetIcon(icon string) {
	m.icon = icon
}

// Setter method for the field desc of type string in the object MenuBuilder
func (m *MenuBuilder) SetDesc(desc string) {
	m.desc = desc
}

// Setter method for the field isActive of type int8 in the object MenuBuilder
func (m *MenuBuilder) SetIsActive(isActive int8) {
	m.isActive = isActive
}
