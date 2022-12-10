package handlers

import (
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeUserHandler(logger *zap.Logger, eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateUser())
	rg.DELETE("/:id", eps.DeleteUser())
	rg.GET("", eps.ListUser())
	rg.PUT("/status", eps.DisableUser())
	// rg.GET("/:id", eps.ViewAccount())
	return rg
}
