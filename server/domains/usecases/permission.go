package usecases

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/domains/entities/mappers"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
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
	mBuild.SetSlug(slug.Make(request.Name))
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
	if err = request.Validate(); err != nil {
		return nil, err
	}
	if err = p.repo.GetPermission().Delete(ctx, request.Id.(uuid.UUID)); err != nil {
		return nil, err
	}
	return new(responses.Empty), nil
}

// ListPermission implements Permission
func (p *permission) ListPermission(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Permission], err error) {
	if request.Limit == 0 {
		request.Limit = 25
	}
	responseBuilder := responses.NewListEntityBuilder[entities.Permission]()
	if validation.IsEmpty(request.Keyword) {
		result := p.repo.GetPermission().GetAll(ctx, request.Limit)
		responseBuilder.SetData(mappers.ListPermissionModelToListEntity(result.Value))
		return responseBuilder.Build(), result.Error
	}
	result := p.repo.GetPermission().GetByKeywords(ctx, request.Keyword, request.Limit)
	responseBuilder.SetData(mappers.ListPermissionModelToListEntity(result.Value))
	return responseBuilder.Build(), result.Error
}

// UpdatePermission implements Permission
func (p *permission) UpdatePermission(ctx context.Context, request *requests.UpdatePermission) (response *responses.Empty, err error) {
	if err = request.Validate(); err != nil {
		return nil, err
	}
	slugs := slug.Make(request.Name)
	mBuild := models.NewPermissionBuilder()
	mBuild.SetId(request.Id.(uuid.UUID))
	mBuild.SetName(request.Name)
	mBuild.SetSlug(slugs)
	mBuild.SetUrl(request.Url)
	mBuild.SetMethod(request.Method)
	model := mBuild.Build()
	if err = p.repo.GetPermission().Update(ctx, model); err != nil {
		return nil, err
	}

	return new(responses.Empty), nil
}

// ViewPermission implements Permission
func (p *permission) ViewPermission(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Permission], err error) {
	// if err = request.Validate(); err != nil {
	// 	return nil, err
	// }

	// result := p.repo.GetPermission().GetById(ctx, request.Id.(uuid.UUID))

	// responseBuilder := responses.NewViewEntityBuilder[*entities.Permission]()
	// responseBuilder.SetData(mappers.PermissionModelToEntity(result.Value))
	// return responseBuilder.Build(), result.Error
	return nil, nil

}

func NewPermission(logger *zap.Logger, repo repositories.Repositories) Permission {
	a := new(permission)
	a.logger = logger
	a.repo = repo
	return a
}
