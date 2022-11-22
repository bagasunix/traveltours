package menu

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commod interface {
	Create(ctx context.Context, m *models.Menu) error
	Update(ctx context.Context, m *models.Menu) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Menu])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Menu])
}

type Repository interface {
	base.Repository
	Commod
	Query
}
