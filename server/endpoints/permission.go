package endpoints

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type PermissionEndpoint interface {
	CreatePermession() gin.HandlerFunc
	UpdatePermission() gin.HandlerFunc
	DeletePermission() gin.HandlerFunc
	ListPermission() gin.HandlerFunc
	ViewPermission() gin.HandlerFunc
}

type permissionHandler struct {
	service domains.Service
}

// DeletePermission implements PermissionEndpoint
func (p *permissionHandler) DeletePermission() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := p.service.DeletePermission(g, req.(*requests.EntityId))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusNoContent, dataRole)
	}
}

// ListPermission implements PermissionEndpoint
func (p *permissionHandler) ListPermission() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeBaseListEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := p.service.ListPermission(g, req.(*requests.BaseList))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusOK, dataRole)
	}
}

// UpdatePermission implements PermissionEndpoint
func (p *permissionHandler) UpdatePermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.UpdatePermission

		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		req.Id = uuid.FromStringOrNil(req.Id.(string))
		dataRole, err := p.service.UpdatePermission(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusOK, dataRole)
	}
}

// ViewPermission implements PermissionEndpoint
func (p *permissionHandler) ViewPermission() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := p.service.ViewPermission(g, req.(*requests.EntityId))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusOK, dataRole)
	}
}

// CreatePermession implements PermissionEndpoint
func (p *permissionHandler) CreatePermession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreatePermission
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		dataAccount, err := p.service.CreatePermission(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusCreated, dataAccount)
	}
}

func NewPermissionEndpoint(svc domains.Service) PermissionEndpoint {
	return &permissionHandler{service: svc}
}
