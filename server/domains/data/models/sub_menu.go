package models

import "github.com/gofrs/uuid"

type SubMenu struct {
	BaseModel
	MenuId   uuid.UUID
	Menu     *Menu `gorm:"foreignKey:MenuId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	Title    string
	Url      string
	Desc     string
	Position int8
	IsActive int8
}

// Builder Object for SubMenu
type SubMenuBuilder struct {
	BaseModelBuilder
	menuId   uuid.UUID
	menu     *Menu
	title    string
	url      string
	desc     string
	position int8
	isActive int8
}

// Constructor for SubMenuBuilder
func NewSubMenuBuilder() *SubMenuBuilder {
	o := new(SubMenuBuilder)
	return o
}

// Build Method which creates SubMenu
func (b *SubMenuBuilder) Build() *SubMenu {
	o := new(SubMenu)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.MenuId = b.menuId
	o.Menu = b.menu
	o.Title = b.title
	o.Url = b.url
	o.Desc = b.desc
	o.Position = b.position
	o.IsActive = b.isActive
	return o
}

// Setter method for the field menuId of type uuid.UUID in the object SubMenuBuilder
func (s *SubMenuBuilder) SetMenuId(menuId uuid.UUID) {
	s.menuId = menuId
}

// Setter method for the field menu of type *Menu in the object SubMenuBuilder
func (s *SubMenuBuilder) SetMenu(menu *Menu) {
	s.menu = menu
}

// Setter method for the field title of type string in the object SubMenuBuilder
func (s *SubMenuBuilder) SetTitle(title string) {
	s.title = title
}

// Setter method for the field url of type string in the object SubMenuBuilder
func (s *SubMenuBuilder) SetUrl(url string) {
	s.url = url
}

// Setter method for the field desc of type string in the object SubMenuBuilder
func (s *SubMenuBuilder) SetDesc(desc string) {
	s.desc = desc
}

// Setter method for the field position of type int8 in the object SubMenuBuilder
func (s *SubMenuBuilder) SetPosition(position int8) {
	s.position = position
}

// Setter method for the field isActive of type int8 in the object SubMenuBuilder
func (s *SubMenuBuilder) SetIsActive(isActive int8) {
	s.isActive = isActive
}
