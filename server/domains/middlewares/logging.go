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

// ViewRole implements domains.Service
func (l *loggingMiddleware) ViewRole(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Role], err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewRole(ctx, request)
}

// ViewRoleById implements domains.Service
func (l *loggingMiddleware) ViewRoleById(ctx context.Context, id uuid.UUID) (role *entities.Role, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewRoleById"), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewRoleById(ctx, id)
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
func (l *loggingMiddleware) AssignPermissionsToRole(ctx context.Context, request *requests.AssignPermissionsToRole) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "AssignPermissionsToRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.AssignPermissionsToRole(ctx, request)
}

// CreateRole implements domains.Service
func (l *loggingMiddleware) CreateRole(ctx context.Context, request *requests.CreateRole) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreateRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateRole(ctx, request)
}

// DeleteRole implements domains.Service
func (l *loggingMiddleware) DeleteRole(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DeleteRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeleteRole(ctx, request)
}

// ListRole implements domains.Service
func (l *loggingMiddleware) ListRole(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Role], err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ListRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ListRole(ctx, request)
}

// RemovePermissionsFromRole implements domains.Service
func (l *loggingMiddleware) RemovePermissionsFromRole(ctx context.Context, request *requests.RemovePermissionsFromRole) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "RemovePermissionsFromRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.RemovePermissionsFromRole(ctx, request)
}

// UpdateRole implements domains.Service
func (l *loggingMiddleware) UpdateRole(ctx context.Context, request *requests.UpdateRole) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "UpdateRole"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdateRole(ctx, request)
}

// CreatePermission implements domains.Service
func (l *loggingMiddleware) CreatePermission(ctx context.Context, request *requests.CreatePermission) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "CreatePermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreatePermission(ctx, request)
}

// DeletePermission implements domains.Service
func (l *loggingMiddleware) DeletePermission(ctx context.Context, request *requests.EntityId) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "DeletePermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeletePermission(ctx, request)
}

// ListPermission implements domains.Service
func (l *loggingMiddleware) ListPermission(ctx context.Context, request *requests.BaseList) (response *responses.ListEntity[entities.Permission], err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ListPermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ListPermission(ctx, request)
}

// UpdatePermission implements domains.Service
func (l *loggingMiddleware) UpdatePermission(ctx context.Context, request *requests.UpdatePermission) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "UpdatePermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdatePermission(ctx, request)
}

// ViewPermission implements domains.Service
func (l *loggingMiddleware) ViewPermission(ctx context.Context, request *requests.EntityId) (response *responses.ViewEntity[*entities.Permission], err error) {
	defer func(begin time.Time) {
		l.logger.Log(zap.InfoLevel, "Middleware Domain", zap.String("method", "ViewPermission"), zap.Any("request", string(request.ToJSON())), zap.Any("response", string(response.ToJSON())), zap.Any("err", err), zap.Any("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewPermission(ctx, request)
}

func LoggingMiddleware(logger *zap.Logger) domains.Middleware {
	return func(next domains.Service) domains.Service {
		return &loggingMiddleware{logger: logger, next: next}
	}
}
