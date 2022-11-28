package handlers

import (
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeRoleHandler(logger *zap.Logger, eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateRole())
	// rg.POST("", eps.CreateUser())
	// rg.Use(middlewares.Auth(logs), middlewares.Permission("admin"))
	// rg.GET("", eps.ListAccount())
	// rg.GET("/:id", eps.ViewAccount())
	// rg.DELETE("/:id", eps.DeleteAccount())
	// rg.PUT("", eps.DisableAccount())
	return rg
}
