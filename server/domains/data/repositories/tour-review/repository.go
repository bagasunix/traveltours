package tourreview

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.TourReview) error
	Update(ctx context.Context, m *models.TourReview) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetByAll(ctx context.Context) (result models.SliceResult[models.TourReview])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.TourReview])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
