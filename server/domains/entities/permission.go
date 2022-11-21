package entities

type Permission struct {
	Entity
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Method string `json:"method"`
	Url    string `json:"uri"`
}

// PermissionBuilder Builder Object for Permission
type PermissionBuilder struct {
	EntityBuilder
	name   string
	slug   string
	url    string
	method string
}

// NewPermissionBuilder Constructor for PermissionBuilder
func NewPermissionBuilder() *PermissionBuilder {
	o := new(PermissionBuilder)
	return o
}

// Build Method which creates Permission
func (p *PermissionBuilder) Build() *Permission {
	o := new(Permission)
	o.Entity = *p.EntityBuilder.Build()
	o.Name = p.name
	o.Slug = p.slug
	o.Method = p.method
	o.Url = p.url
	return o
}

// SetName Setter method for the field name of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetName(name string) *PermissionBuilder {
	p.name = name
	return p
}

// SetUri Setter method for the field uri of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetUrl(url string) *PermissionBuilder {
	p.url = url
	return p
}

// SetMethod Setter method for the field method of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetMethod(method string) *PermissionBuilder {
	p.method = method
	return p
}

// Setter method for the field slug of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetSlug(slug string) {
	p.slug = slug
}
