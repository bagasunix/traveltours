package handlers

import (
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakePermissionHandler(logger *zap.Logger, eps endpoints.PermissionEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreatePermession())
	return rg
}
