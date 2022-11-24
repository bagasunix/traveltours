package inits

import (
	"net/http"

	"github.com/bagasunix/traveltours/server/endpoints"
	transportHttp "github.com/bagasunix/traveltours/server/transports/http"
	"go.uber.org/zap"
)

func InitHttpHandler(logger *zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	return transportHttp.NewHttpHandler(logger, endpoints)
}
