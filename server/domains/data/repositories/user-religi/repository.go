package userreligi

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.UserReligi) error
	Update(ctx context.Context, m *models.UserReligi) error
	Delete(ctx context.Context, id uuid.UUID) error
}
type Query interface {
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserReligi])
	GetAll(ctx context.Context) (result models.SliceResult[models.UserReligi])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
