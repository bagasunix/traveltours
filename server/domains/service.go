package domains

import (
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"go.uber.org/zap"
)

type Service interface {
}

type service struct {
}

type ServiceBuild struct {
	logger *zap.Logger
	repo   repositories.Repositories
	mdw    []Middleware
}

func NewServiceBuild(logger *zap.Logger, repo repositories.Repositories) *ServiceBuild {
	a := new(ServiceBuild)
	a.logger = logger
	a.repo = repo
	return a
}

func buildService(logger *zap.Logger, repo repositories.Repositories) Service {
	svc := new(service)
	return svc
}

func (s *ServiceBuild) Build() Service {
	svc := buildService(s.logger, s.repo)
	for _, m := range s.mdw {
		svc = m(svc)
	}
	return svc
}

// Setter method for the field mdw of type []Middleware in the object ServiceBuild
func (s *ServiceBuild) SetMdw(mdw []Middleware) {
	s.mdw = mdw
}
