package rolepermission

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, m *models.RolePermission) error
	Update(ctx context.Context, m *models.RolePermission) error
	Delete(ctx context.Context, id uuid.UUID) error
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
