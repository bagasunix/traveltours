package inits

import (
	"github.com/bagasunix/traveltours/pkg/env"
	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/middlewares"
	"go.uber.org/zap"
)

func InitService(logger *zap.Logger, configs *env.Configs, repo repositories.Repositories) domains.Service {
	serviceBuilder := domains.NewServiceBuild(logger, repo)
	serviceBuilder.SetMdw(getServiceMiddleware(logger))
	return serviceBuilder.Build()
}

func getServiceMiddleware(logger *zap.Logger) []domains.Middleware {
	var mw []domains.Middleware
	mw = addDefaultServiceMiddleware(logger, mw)
	return mw
}

func addDefaultServiceMiddleware(logger *zap.Logger, mw []domains.Middleware) []domains.Middleware {
	return append(mw, middlewares.LoggingMiddleware(logger))
}
