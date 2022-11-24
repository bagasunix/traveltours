package middlewares

import (
	"time"

	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logging(logger *zap.Logger) endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log(zap.InfoLevel, err.Error(), zap.Any("took", time.Since(begin)))
			}(time.Now())
			return e(ctx, request)
		}
	}
}
