package endpoints

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type RoleEndpoint interface {
	CreateRole() gin.HandlerFunc
	UpdateRole() gin.HandlerFunc
	GetAllRole() gin.HandlerFunc
	DeleteRole() gin.HandlerFunc
}

type roleHandler struct {
	service domains.Service
}

// DeleteRole implements RoleEndpoint
func (r *roleHandler) DeleteRole() gin.HandlerFunc {
	return func(g *gin.Context) {
		req, err := decodeByEntityIdEndpoint(g)
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		dataRole, err := r.service.DeleteRole(g, req.(*requests.EntityId))
		if err != nil {
			utils.EncodeError(g, err, g.Writer)
			return
		}
		g.JSON(http.StatusNoContent, dataRole)
	}
}

// GetAllRole implements RoleEndpoint
func (r *roleHandler) GetAllRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.BaseList
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		dataRole, err := r.service.ListRole(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusOK, dataRole)
	}
}

// UpdateRole implements RoleEndpoint
func (r *roleHandler) UpdateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.UpdateRole
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		req.Id = uuid.FromStringOrNil(req.Id.(string))
		dataRole, err := r.service.UpdateRole(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusOK, dataRole)
	}
}

// CreateRole implements RoleEndpoint
func (r *roleHandler) CreateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateRole
		if err := ctx.Bind(&req); err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		dataRole, err := r.service.CreateRole(ctx, &req)
		if err != nil {
			utils.EncodeError(ctx, err, ctx.Writer)
			return
		}
		ctx.JSON(http.StatusCreated, dataRole)
	}
}

func NewRoleEndpoint(svc domains.Service) RoleEndpoint {
	return &roleHandler{service: svc}
}
