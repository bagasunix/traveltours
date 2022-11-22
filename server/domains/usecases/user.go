package usecases

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/repositories"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
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
	repo   repositories.Repositories
	logger *zap.Logger
}

// CreateUser implements User
func (*user) CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, err error) {
	panic("unimplemented")
}

// DeleteUser implements User
func (*user) DeleteUser(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// DisableAccount implements User
func (*user) DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error) {
	panic("unimplemented")
}

// ListUser implements User
func (*user) ListUser(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.User], err error) {
	panic("unimplemented")
}

// ViewUser implements User
func (*user) ViewUser(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error) {
	panic("unimplemented")
}

// ListUserByLimit implements User
func (*user) ListUserByLimit(ctx context.Context, limit int64) (users []entities.User, err error) {
	panic("unimplemented")
}

// ViewUserByEmail implements User
func (*user) ViewUserByEmail(ctx context.Context, emailOrMsisdn string) (user *entities.User, err error) {
	panic("unimplemented")
}

// ViewUserById implements User
func (*user) ViewUserById(ctx context.Context, id uuid.UUID) (user *entities.User, err error) {
	panic("unimplemented")
}

func NewAccount(logger *zap.Logger, repo repositories.Repositories) User {
	a := new(user)
	a.logger = logger
	a.repo = repo
	return a
}
