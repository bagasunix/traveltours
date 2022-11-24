package models

type Permission struct {
	BaseModel
	Name   string `gorm:"uniqueIndex;size:200"`
	Slug   string
	Method string `gorm:"uniqueIndex:permission_unique_index"`
	Url    string `gorm:"uniqueIndex:permission_unique_index"`
	Roles  []Role `gorm:"many2many:role_permissions;"`
}

// Builder Object for Permission
type PermissionBuilder struct {
	BaseModelBuilder
	name   string
	slug   string
	method string
	url    string
}

// Constructor for PermissionBuilder
func NewPermissionBuilder() *PermissionBuilder {
	o := new(PermissionBuilder)
	return o
}

// Build Method which creates Permission
func (b *PermissionBuilder) Build() *Permission {
	o := new(Permission)
	o.BaseModel = *b.BaseModelBuilder.Build()
	o.Name = b.name
	o.Slug = b.slug
	o.Method = b.method
	o.Url = b.url
	return o
}

// Setter method for the field method of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetMethod(method string) {
	p.method = method
}

// Setter method for the field url of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetUrl(url string) {
	p.url = url
}

// Setter method for the field name of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetName(name string) {
	p.name = name
}

// Setter method for the field slug of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetSlug(slug string) {
	p.slug = slug
}
