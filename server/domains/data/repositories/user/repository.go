package user

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.User) error
	Update(ctx context.Context, m *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteTx(ctx context.Context, tx any, id uuid.UUID) error
}

type Query interface {
	GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.User])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.User])
	GetByKeywordEmail(ctx context.Context, entity string, limit int64) (result models.SliceResult[models.User])
	UpdateByStatus(ctx context.Context, m *models.User) error
}

type Repository interface {
	base.Repository
	Commond
	Query
}
