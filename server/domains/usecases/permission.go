package usecases

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	"go.uber.org/zap"
)

type Permission interface {
	CreatePermission(ctx context.Context, request *requests.CreatePermission) (response *responses.EntityId, err error)
	UpdatePermission(ctx context.Context, request *requests.UpdatePermission) (response *responses.Empty, err error)
	DeletePermission(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error)
	ListPermission(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Permission], err error)
	ViewPermission(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Permission], err error)
}

type permission struct {
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreatePermission implements Permission
func (p *permission) CreatePermission(ctx context.Context, request *requests.CreatePermission) (response *responses.EntityId, err error) {
	if err = request.Validate(); err != nil {
		return nil, err
	}
	mBuild := models.NewPermissionBuilder()
	mBuild.SetId(helpers.GenerateUUIDV1(p.logger))
	mBuild.SetName(request.Name)
	mBuild.SetSlug(request.Slug)
	mBuild.SetUrl(request.Url)
	mBuild.SetMethod(request.Method)
	if err = p.repo.GetPermission().Create(ctx, mBuild.Build()); err != nil {
		return nil, err
	}
	responseBuilder := responses.NewEntityIdBuilder()
	return responseBuilder.SetId(mBuild.Build().Id).Build(), nil
}

// DeletePermission implements Permission
func (p *permission) DeletePermission(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ListPermission implements Permission
func (p *permission) ListPermission(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Permission], err error) {
	panic("unimplemented")
}

// UpdatePermission implements Permission
func (p *permission) UpdatePermission(ctx context.Context, request *requests.UpdatePermission) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ViewPermission implements Permission
func (p *permission) ViewPermission(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Permission], err error) {
	panic("unimplemented")
}

func NewPermission(logger *zap.Logger, repo repositories.Repositories) Permission {
	a := new(permission)
	a.logger = logger
	a.repo = repo
	return a
}
