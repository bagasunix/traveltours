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
}

type roleHandler struct {
	service domains.Service
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
