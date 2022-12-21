package usecases

import (
	"context"
	"strings"
	"time"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/domains/entities/mappers"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserWithEndpoint interface {
	CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, err error)
	DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error)
	DeleteUser(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error)
	ListUser(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.User], err error)
	ViewUser(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error)
}

type UserWithNoEndpoint interface {
	ListUserByLimit(ctx context.Context, limit int64) (users []entities.User, err error)
	ViewUserByEmail(ctx context.Context, emailOrMsisdn string) (user *entities.User, err error)
	ViewUserById(ctx context.Context, id uuid.UUID) (user *entities.User, err error)
}

type User interface {
	UserWithEndpoint
	UserWithNoEndpoint
}

type user struct {
	c      *gin.Context
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreateUser implements User
func (u *user) CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, err error) {
	resBuilder := responses.NewEntityIdBuilder()

	if err = request.Validate(); err != nil {
		return resBuilder.Build(), err
	}

	if helpers.IsEmailValid(request.Email) != true {
		return resBuilder.Build(), errors.ErrValidEmail(u.logger, request.Email)
	}

	if err = u.repo.GetRole().GetById(ctx, request.RoleId).Error; err != nil {
		return resBuilder.Build(), err
	}

	mUser := models.NewUserBuilder()
	mUser.SetId(helpers.GenerateUUIDV1(u.logger))
	mUser.SetEmail(request.Email)
	mUser.SetPassword(helpers.HashAndSalt([]byte(request.Password)))
	mUser.SetRoleId(request.RoleId)
	mUser.SetStatusId(1)

	if err = u.repo.GetUser().Create(ctx, mUser.Build()); err != nil {
		return resBuilder.Build(), err
	}

	return resBuilder.SetId(mUser.Build().Id).Build(), nil

}

// DeleteUser implements User
func (u *user) DeleteUser(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	if err = u.repo.GetUser().GetById(ctx, request.Id.(uuid.UUID)).Error; err != nil {
		return nil, err
	}

	if err = u.repo.GetUserDetail().GetByUserId(ctx, request.Id.(uuid.UUID)).Error; err != nil {
		if err = u.repo.GetUser().GetById(ctx, request.Id.(uuid.UUID)).Error; err != nil {
			return nil, err
		}
	}

	tx := u.repo.GetUser().GetConnection().(*gorm.DB).Begin()
	if err = u.repo.GetUser().DeleteTx(ctx, tx, request.Id.(uuid.UUID)); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = u.repo.GetUserDetail().DeleteTx(ctx, tx, request.Id.(uuid.UUID)); err != nil {
		tx.Rollback()
		return nil, err
	}

	return new(responses.Empty), errors.ErrSomethingWrong(u.logger, tx.Commit().Error)
}

// DisableAccount implements User
func (u *user) DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	uUid := uuid.FromStringOrNil(req.Id.(string))
	result := u.repo.GetUser().GetById(ctx, uUid)
	if result.Error != nil {
		return nil, result.Error
	}
	if err = u.repo.GetUserStatus().GetById(ctx, req.StatusId).Error; err != nil {
		return nil, err
	}
	mBuild := models.NewUserBuilder()
	mBuild.SetId(uUid)
	mBuild.SetStatusId(req.StatusId)
	mBuild.SetUpdatedAt(time.Now())
	return new(responses.Empty), u.repo.GetUser().UpdateByStatus(ctx, mBuild.Build())
}

// ListUser implements User
func (u *user) ListUser(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.User], err error) {
	if request.Limit == 0 {
		request.Limit = 25
	}
	resBuild := responses.NewListEntityBuilder[entities.User]()
	var entitiesUser []entities.User
	if validation.IsEmpty(request.Keyword) {
		result := u.repo.GetUser().GetByAll(ctx, request.Limit)
		for _, v := range result.Value {
			resRole := u.repo.GetRole().GetById(ctx, v.RoleId)
			resStatus := u.repo.GetUserStatus().GetById(ctx, v.StatusId)

			newEntitiesUser := entities.NewUserBuilder()
			newEntitiesUser.SetId(v.Id)
			newEntitiesUser.SetEmail(v.Email)
			newEntitiesUser.SetRole(strings.ToLower(resRole.Value.Name))
			newEntitiesUser.SetStatus(strings.ToLower(resStatus.Value.Name))
			newEntitiesUser.SetCreatedAt(v.CreatedAt)
			newEntitiesUser.SetCreatedBy(v.CreatedBy)
			entitiesUser = append(entitiesUser, *newEntitiesUser.Build())

		}
		resBuild.SetData(entitiesUser)
		return resBuild.Build(), result.Error
	}
	result := u.repo.GetUser().GetByKeywordEmail(ctx, request.Keyword, request.Limit)
	resBuild.SetData(mappers.ListUserModelToListEntity(result.Value))
	return resBuild.Build(), result.Error
}

// ViewUser implements User
func (u *user) ViewUser(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error) {
	panic("unimplemented")
}

// ListUserByLimit implements User
func (u *user) ListUserByLimit(ctx context.Context, limit int64) (users []entities.User, err error) {
	panic("unimplemented")
}

// ViewUserByEmail implements User
func (u *user) ViewUserByEmail(ctx context.Context, emailOrMsisdn string) (user *entities.User, err error) {
	panic("unimplemented")
}

// ViewUserById implements User
func (u *user) ViewUserById(ctx context.Context, id uuid.UUID) (user *entities.User, err error) {
	panic("unimplemented")
}

func NewUser(logger *zap.Logger, repo repositories.Repositories) User {
	a := new(user)
	a.logger = logger
	a.repo = repo
	return a
}
