package models

type Permission struct {
	BaseModel
	Name   string `gorm:"size:50"`
	Slug   string
	Method string
	Url    []string
}

// Builder Object for Permission
type PermissionBuilder struct {
	BaseModelBuilder
	name   string
	slug   string
	method string
	url    []string
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

// Getter method for the field slug of type string in the object PermissionBuilder
func (p *PermissionBuilder) Slug() string {
	return p.slug
}

// Setter method for the field method of type []string in the object PermissionBuilder
func (p *PermissionBuilder) SetMethod(method string) {
	p.method = method
}

// Setter method for the field url of type []string in the object PermissionBuilder
func (p *PermissionBuilder) SetUrl(url []string) {
	p.url = url
}
