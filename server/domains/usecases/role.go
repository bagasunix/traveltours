package usecases

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type RoleWithEndpoint interface {
	CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error)
	UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error)
	DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error)
	ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error)
	ViewRole(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Role], err error)
}

type RoleWithNoEndpoint interface {
	ViewRoleById(ctx context.Context, Id uuid.UUID) (role *entities.Role, err error)
	ViewRoleBySelectedId(ctx context.Context, ids []uuid.UUID) (roles []entities.Role, err error)
}

type Role interface {
	RoleWithEndpoint
	RoleWithNoEndpoint
}

type role struct {
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreateRole implements Role
func (r *role) CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error) {
	if err = request.Validate(); err != nil {
		return nil, err
	}

	roleMBuild := models.NewRoleBuilder()
	roleMBuild.SetId(helpers.GenerateUUIDV1(r.logger))
	roleMBuild.SetName(request.Name)
	roleMBuild.SetGroup(request.Group)

	if err = r.repo.GetRole().Create(ctx, roleMBuild.Build()); err != nil {
		return nil, err
	}

	resBuild := responses.NewEntityIdBuilder()
	resBuild.SetId(roleMBuild.Build())
	return resBuild.Build(), nil
}

// DeleteRole implements Role
func (r *role) DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ListRole implements Role
func (r *role) ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error) {
	panic("unimplemented")
}

// UpdateRole implements Role
func (r *role) UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ViewRole implements Role
func (r *role) ViewRole(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Role], err error) {
	panic("unimplemented")
}

// ViewRoleById implements Role
func (r *role) ViewRoleById(ctx context.Context, Id uuid.UUID) (role *entities.Role, err error) {
	panic("unimplemented")
}

// ViewRoleBySelectedId implements Role
func (r *role) ViewRoleBySelectedId(ctx context.Context, ids []uuid.UUID) (roles []entities.Role, err error) {
	panic("unimplemented")
}

func NewRole(logger *zap.Logger, repo repositories.Repositories) Role {
	a := new(role)
	a.logger = logger
	a.repo = repo
	return a
}
