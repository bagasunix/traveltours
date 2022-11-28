package endpoints

import (
	"github.com/bagasunix/traveltours/server/domains"
	"github.com/gin-gonic/gin"
)

type Endpoint func(ctx *gin.Context) (res interface{}, err error)
type Middleware func(Endpoint) Endpoint

type Endpoints struct {
	UserEndpoint       UserEndpoint
	PermissionEndpoint PermissionEndpoint
	RoleEndpoint       RoleEndpoint
}

// Builder Object for Endpoints
type EndpointsBuilder struct {
	service domains.Service
	mdw     map[string][]Middleware
}

// Constructor for EndpointsBuilder
func NewEndpointsBuilder() *EndpointsBuilder {
	o := new(EndpointsBuilder)
	return o
}

// Build Method which creates Endpoints
func (b *EndpointsBuilder) Build() Endpoints {
	o := new(Endpoints)
	o.UserEndpoint = NewUserEndpoint(b.service)
	o.PermissionEndpoint = NewPermissionEndpoint(b.service)
	o.RoleEndpoint = NewRoleEndpoint(b.service)
	return *o
}

// Setter method for the field service of type domains.Service in the object EndpointsBuilder
func (e *EndpointsBuilder) SetService(service domains.Service) {
	e.service = service
}

// Setter method for the field mds of type gin.HandlerFunc in the object EndpointsBuilder
func (e *EndpointsBuilder) SetMdw(mdw map[string][]Middleware) {
	e.mdw = mdw
}
