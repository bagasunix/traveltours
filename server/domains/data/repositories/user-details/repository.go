package userdetails

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.UserDetails) error
	Update(ctx context.Context, m *models.UserDetails) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteTx(ctx context.Context, tx any, id uuid.UUID) error
}

type Query interface {
	GetByAll(ctx context.Context) (result models.SliceResult[models.UserDetails])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserDetails])
	GetByUserId(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserDetails])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
