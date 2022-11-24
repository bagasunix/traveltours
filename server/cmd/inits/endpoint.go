package inits

import (
	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/endpoints"
	"github.com/bagasunix/traveltours/server/endpoints/middlewares"
	"go.uber.org/zap"
)

func InitEndpoints(logger *zap.Logger, s domains.Service) endpoints.Endpoints {
	a := endpoints.NewEndpointsBuilder()
	a.SetMdw(getEndpointMiddleware(logger))
	a.SetService(s)
	return a.Build()
}

func getEndpointMiddleware(logger *zap.Logger) (mw map[string][]endpoints.Middleware) {
	mw = map[string][]endpoints.Middleware{}
	addDefaultEndpointMiddleware(logger, mw)
	return mw
}

func middlewaresWithAuthentication(logger *zap.Logger, method string) []endpoints.Middleware {
	mw := defaultMiddlewares(logger, method)
	return mw
	// return append(mw, middlewares.Authentication())
}

func defaultMiddlewares(logger *zap.Logger, method string) []endpoints.Middleware {
	return []endpoints.Middleware{
		middlewares.Logging(logger.With(zap.Any("method", method))),
		// middlewares.Logging(log.With(logger, "method", method)),
	}
}

func addDefaultEndpointMiddleware(logger *zap.Logger, mw map[string][]endpoints.Middleware) {
	mw[endpoints.CREATE_USER] = middlewaresWithAuthentication(logger, endpoints.CREATE_USER)
}
