package rolepermission

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, m *models.RolePermission) error
	CreateTx(ctx context.Context, tx any, model *models.RolePermission) error
	CreateTxBatch(ctx context.Context, tx any, models []models.RolePermission) error
	CreateBatch(ctx context.Context, models []models.RolePermission) error
	Update(ctx context.Context, m *models.RolePermission) error
	Delete(ctx context.Context, roleId uuid.UUID, permissionId uuid.UUID) error
	DeleteBatch(ctx context.Context, roleId uuid.UUID, permissionIds []uuid.UUID) error
}

type Query interface {
	GetByRoleId(ctx context.Context, roleId uuid.UUID) (result models.SliceResult[*models.RolePermission])
	GetBySelectedRoleId(ctx context.Context, roleIds []uuid.UUID) (result models.SliceResult[models.RolePermission])
	GetById(ctx context.Context, id uuid.UUID) (result models.SliceResult[models.RolePermission])
}

type Repository interface {
	base.Repository
	Query
	Command
}
