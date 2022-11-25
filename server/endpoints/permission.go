package endpoints

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/utils"
	"github.com/gin-gonic/gin"
)

type PermissionEndpoint interface {
	CreatePermession() gin.HandlerFunc
}

type permissionHandler struct {
	service domains.Service
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
