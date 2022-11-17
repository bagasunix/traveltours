package tour

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.Tour) error
	Update(ctx context.Context, m *models.Tour) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context) (result models.SliceResult[models.Tour])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Tour])
	GetByKeyword(ctx context.Context, entity string) (result models.SliceResult[models.Tour])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
