package endpoints

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type RoleEndpoint interface {
	CreateRole() gin.HandlerFunc
}

type roleHandler struct {
	service domains.Service
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
