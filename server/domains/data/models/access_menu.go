package models

import "github.com/gofrs/uuid"

type AccessMenu struct {
	BaseModel
	GroupMenuId uuid.UUID
	GroupMenu   *GroupMenu `gorm:"foreignKey:GroupMenuId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	MenuId      uuid.UUID
	Menu        *Menu `gorm:"foreignKey:MenuId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	SubMenuId   uuid.UUID
	SubMenu     *SubMenu `gorm:"foreignKey:SubMenuId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
	RoleId      uuid.UUID
	Role        *Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:Restrict"`
}

// Builder Object for AccessMenu
type AccessMenuBuilder struct {
	BaseModelBuilder
	groupMenuId uuid.UUID
	groupMenu   *GroupMenu
	menuId      uuid.UUID
	menu        *Menu
	subMenuId   uuid.UUID
	subMenu     *SubMenu
	roleId      uuid.UUID
	role        *Role
}

// Constructor for AccessMenuBuilder
func NewAccessMenuBuilder() *AccessMenuBuilder {
	o := new(AccessMenuBuilder)
	return o
}

// Build Method which creates AccessMenu
func (b *AccessMenuBuilder) Build() *AccessMenu {
	o := new(AccessMenu)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.GroupMenuId = b.groupMenuId
	o.GroupMenu = b.groupMenu
	o.MenuId = b.menuId
	o.Menu = b.menu
	o.SubMenuId = b.subMenuId
	o.SubMenu = b.subMenu
	o.RoleId = b.roleId
	o.Role = b.role
	return o
}

// Setter method for the field groupMenuId of type uuid.UUID in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetGroupMenuId(groupMenuId uuid.UUID) {
	p.groupMenuId = groupMenuId
}

// Setter method for the field groupMenu of type *GroupMenu in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetGroupMenu(groupMenu *GroupMenu) {
	p.groupMenu = groupMenu
}

// Setter method for the field menuId of type uuid.UUID in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetMenuId(menuId uuid.UUID) {
	p.menuId = menuId
}

// Setter method for the field menu of type *Menu in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetMenu(menu *Menu) {
	p.menu = menu
}

// Setter method for the field subMenuId of type uuid.UUID in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetSubMenuId(subMenuId uuid.UUID) {
	p.subMenuId = subMenuId
}

// Setter method for the field subMenu of type *SubMenu in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetSubMenu(subMenu *SubMenu) {
	p.subMenu = subMenu
}

// Setter method for the field roleId of type uuid.UUID in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetRoleId(roleId uuid.UUID) {
	p.roleId = roleId
}

// Setter method for the field role of type *Role in the object AccessMenuBuilder
func (p *AccessMenuBuilder) SetRole(role *Role) {
	p.role = role
}
