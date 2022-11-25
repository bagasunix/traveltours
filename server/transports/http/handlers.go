package http

import (
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/bagasunix/traveltours/server/transports/http/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewHttpHandler(logger *zap.Logger, eps endpoints.Endpoints) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	// r.Use(ginzap.RecoveryWithZap(logs, true))
	// r.Use(middlewares.GinContextToContextMiddleware())
	// r.Use(middlewares.CORSMiddleware())

	// Create an user group
	handlers.MakeUserHandler(logger, eps.UserEndpoint, r.Group("/v1/user"))
	handlers.MakePermissionHandler(logger, eps.PermissionEndpoint, r.Group("/v1/permission"))
	return r
}
