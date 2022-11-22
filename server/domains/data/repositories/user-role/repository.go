package userrole

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.Role) error
	CreateTx(ctx context.Context, tx any, m *models.Role) error
	Update(ctx context.Context, m models.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.Role])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Role])
	GetByGroup(ctx context.Context, group string) (result models.SliceResult[models.Role])
	GetByKeywords(ctx context.Context, entity string, limit int64) (result models.SliceResult[models.Role])
	GetBySelectedId(ctx context.Context, ids []uuid.UUID) (result models.SliceResult[models.Role])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
