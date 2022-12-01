package handlers

import (
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeRoleHandler(logger *zap.Logger, eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateRole())
	rg.PUT("", eps.UpdateRole())
	rg.GET("", eps.GetAllRole())
	rg.DELETE("/:id", eps.DeleteRole())
	rg.GET("/:id", eps.ViewRole())
	rg.PATCH("/:id/permissions", eps.AssignPermissionsToRole())
	rg.DELETE("/:id/permissions", eps.RemovePermissionsFromRole())
	return rg
}
