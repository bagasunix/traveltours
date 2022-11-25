package middlewares

import (
	"context"
	"time"

	"github.com/bagasunix/traveltours/server/domains"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/bagasunix/traveltours/server/endpoints/requests"
	"github.com/bagasunix/traveltours/server/endpoints/responses"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type loggingMiddleware struct {
	logger *zap.Logger
	next   domains.Service
}

// CreateUser implements domains.Service
func (l *loggingMiddleware) CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreateUser"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateUser(ctx, request)
}

// DeleteUser implements domains.Service
func (*loggingMiddleware) DeleteUser(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// DisableAccount implements domains.Service
func (*loggingMiddleware) DisableAccount(ctx context.Context, req *requests.DisableAccount) (res *responses.Empty, err error) {
	panic("unimplemented")
}

// ListUser implements domains.Service
func (*loggingMiddleware) ListUser(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.User], err error) {
	panic("unimplemented")
}

// ViewUser implements domains.Service
func (*loggingMiddleware) ViewUser(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error) {
	panic("unimplemented")
}

// ListUserByLimit implements domains.Service
func (*loggingMiddleware) ListUserByLimit(ctx context.Context, limit int64) (users []entities.User, err error) {
	panic("unimplemented")
}

// ViewUserByEmail implements domains.Service
func (*loggingMiddleware) ViewUserByEmail(ctx context.Context, emailOrMsisdn string) (user *entities.User, err error) {
	panic("unimplemented")
}

// ViewUserById implements domains.Service
func (*loggingMiddleware) ViewUserById(ctx context.Context, id uuid.UUID) (user *entities.User, err error) {
	panic("unimplemented")
}

// AssignPermissionsToRole implements domains.Service
func (*loggingMiddleware) AssignPermissionsToRole(ctx context.Context, request *requests.AssignPermissionsToRole) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// CreateRole implements domains.Service
func (*loggingMiddleware) CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error) {
	panic("unimplemented")
}

// DeleteRole implements domains.Service
func (*loggingMiddleware) DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ListRole implements domains.Service
func (*loggingMiddleware) ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error) {
	panic("unimplemented")
}

// RemovePermissionsFromRole implements domains.Service
func (*loggingMiddleware) RemovePermissionsFromRole(ctx context.Context, request *requests.RemovePermissionsFromRole) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// UpdateRole implements domains.Service
func (*loggingMiddleware) UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ViewRoleById implements domains.Service
func (*loggingMiddleware) ViewRoleById(ctx context.Context, id uuid.UUID) (response *responses.ViewEntity[*entities.Role], err error) {
	panic("unimplemented")
}

// CreatePermission implements domains.Service
func (l *loggingMiddleware) CreatePermission(ctx context.Context, request *requests.CreatePermission) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreatePermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreatePermission(ctx, request)
}

// DeletePermission implements domains.Service
func (*loggingMiddleware) DeletePermission(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ListPermission implements domains.Service
func (*loggingMiddleware) ListPermission(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Permission], err error) {
	panic("unimplemented")
}

// UpdatePermission implements domains.Service
func (*loggingMiddleware) UpdatePermission(ctx context.Context, request *requests.UpdatePermission) (response *responses.Empty, err error) {
	panic("unimplemented")
}

// ViewPermission implements domains.Service
func (*loggingMiddleware) ViewPermission(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Permission], err error) {
	panic("unimplemented")
}

func LoggingMiddleware(logger *zap.Logger) domains.Middleware {
	return func(next domains.Service) domains.Service {
		return &loggingMiddleware{logger: logger, next: next}
	}
}
