package usecases

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/domains/entities/mappers"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleWithEndpoint interface {
	CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error)
	UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error)
	DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error)
	ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error)
	AssignPermissionsToRole(ctx context.Context, request *requests.AssignPermissionsToRole) (response *responses.Empty, err error)
	RemovePermissionsFromRole(ctx context.Context, request *requests.RemovePermissionsFromRole) (response *responses.Empty, err error)
}

type RoleWithNoEndpoint interface {
	ViewRoleById(ctx context.Context, id uuid.UUID) (response *responses.ViewEntity[*entities.Role], err error)
}

type Role interface {
	RoleWithEndpoint
	RoleWithNoEndpoint
}

type role struct {
	repo   repositories.Repositories
	logger *zap.Logger
}

// AssignPermissionsToRole implements Role
func (r *role) AssignPermissionsToRole(ctx context.Context, request *requests.AssignPermissionsToRole) (response *responses.Empty, err error) {
	rolePermissionModelBuilder := models.NewRolePermissionBuilder()
	var rolePermissionModels []models.RolePermission
	for _, v := range request.PermissionIds {
		rolePermissionModelBuilder.SetRoleId(request.RoleId.(uuid.UUID))
		rolePermissionModelBuilder.SetPermissionId(v.(uuid.UUID))
		rolePermissionModels = append(rolePermissionModels, *rolePermissionModelBuilder.Build())
	}

	if err = r.repo.GetRolePermission().CreateBatch(ctx, rolePermissionModels); err != nil {
		return nil, err
	}
	return new(responses.Empty), nil
}

// RemovePermissionsFromRole implements Role
func (r *role) RemovePermissionsFromRole(ctx context.Context, request *requests.RemovePermissionsFromRole) (response *responses.Empty, err error) {
	var ids []uuid.UUID
	for _, i := range request.PermissionIds {
		ids = append(ids, i.(uuid.UUID))
	}
	if err = r.repo.GetRolePermission().DeleteBatch(ctx, request.RoleId.(uuid.UUID), ids); err != nil {
		return nil, err
	}
	return new(responses.Empty), nil
}

// CreateRole implements Role
func (r *role) CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error) {
	if err = request.Validate(); err != nil {
		return nil, err
	}

	roleModelBuilder := models.NewRoleBuilder()
	roleModelBuilder.SetId(helpers.GenerateUUIDV1(r.logger))
	roleModelBuilder.SetName(request.Name)
	roleModelBuilder.SetGroup(request.Group)
	roleModelBuilder.SetIsActive("1")

	roleModel := roleModelBuilder.Build()

	rolePermissionModelBuilder := models.NewRolePermissionBuilder()
	var rolePermissionModels []models.RolePermission
	for _, i := range request.PermissionIds {
		rolePermissionModelBuilder.SetRoleId(roleModel.Id)
		rolePermissionModelBuilder.SetPermissionId(i)
		rolePermissionModels = append(rolePermissionModels, *rolePermissionModelBuilder.Build())
	}

	tx := r.repo.GetRole().GetConnection().(*gorm.DB).Begin()

	if err = r.repo.GetRole().CreateTx(ctx, tx, roleModel); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = r.repo.GetRolePermission().CreateTxBatch(ctx, tx, rolePermissionModels); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = errors.ErrSomethingWrong(r.logger, tx.Commit().Error); err != nil {
		tx.Rollback()
		return nil, err
	}

	responseEntityIdBuilder := responses.NewEntityIdBuilder()
	responseEntityIdBuilder.SetId(roleModel)
	return responseEntityIdBuilder.Build(), nil
}

// DeleteRole implements Role
func (r *role) DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return new(responses.Empty), r.repo.GetRole().Delete(ctx, request.Id.(uuid.UUID))
}

// ListRole implements Role
func (r *role) ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error) {
	if request.Limit == 0 {
		request.Limit = 25
	}
	resBuild := responses.NewListEntityBuilder[entities.Role]()
	if validation.IsEmpty(request.Keyword) {
		result := r.repo.GetRole().GetByAll(ctx, request.Limit)
		resBuild.SetData(mappers.ListRoleModelToListEntity(result.Value))
		return resBuild.Build(), result.Error
	}
	result := r.repo.GetRole().GetByKeywords(ctx, request.Keyword, request.Limit)
	resBuild.SetData(mappers.ListRoleModelToListEntity(result.Value))
	return resBuild.Build(), result.Error
}

// UpdateRole implements Role
func (r *role) UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error) {
	if err = request.Validate(); err != nil {
		return nil, err
	}

	mRole := models.NewRoleBuilder()
	mRole.SetId(request.Id.(uuid.UUID))
	mRole.SetName(request.Name)
	mRole.SetGroup(request.Group)
	mRole.SetDesc(request.Desc)
	mRole.SetIsActive(request.IsActive)

	return new(responses.Empty), r.repo.GetRole().Update(ctx, *mRole.Build())
}

// ViewRoleById implements Role
func (r *role) ViewRoleById(ctx context.Context, id uuid.UUID) (response *responses.ViewEntity[*entities.Role], err error) {
	if validation.IsEmpty(id) {
		return nil, errors.ErrInvalidAttributes("role id")
	}

	roleResult := r.repo.GetRole().GetById(ctx, id)
	if roleResult.Error != nil {
		return nil, roleResult.Error
	}

	resBuilder := responses.NewViewEntityBuilder[*entities.Role]()
	return resBuilder.SetData(mappers.RoleModelToEntity(roleResult.Value)).Build(), nil
}

func NewRole(logger *zap.Logger, repo repositories.Repositories) Role {
	a := new(role)
	a.logger = logger
	a.repo = repo
	return a
}
