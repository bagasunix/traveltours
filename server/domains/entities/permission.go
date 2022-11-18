package entities

type Permission struct {
	Entity
	Name   string `json:"name"`
	Slug string
	Method string `json:"method"`
	Uri    string `json:"uri"`
}

// PermissionBuilder Builder Object for Permission
type PermissionBuilder struct {
	EntityBuilder
	name   string
	uri    string
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
	o.Method = p.method
	o.Uri = p.uri
	return o
}

// SetName Setter method for the field name of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetName(name string) *PermissionBuilder {
	p.name = name
	return p
}

// SetUri Setter method for the field uri of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetUri(uri string) *PermissionBuilder {
	p.uri = uri
	return p
}

// SetMethod Setter method for the field method of type string in the object PermissionBuilder
func (p *PermissionBuilder) SetMethod(method string) *PermissionBuilder {
	p.method = method
	return p
}
