package permission

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Command interface {
	Create(ctx context.Context, model *models.Permission) error
	Update(ctx context.Context, model *models.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Permission])
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Permission])
	GetByKeywords(ctx context.Context, keyword string, limit int64) (result models.SliceResult[models.Permission])
	GetBySelectedId(ctx context.Context, ids []uuid.UUID) (result models.SliceResult[models.Permission])
}

type Repository interface {
	base.Repository
	Query
	Command
}
