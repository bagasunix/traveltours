package submenu

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
	"github.com/gofrs/uuid"
)

type Commond interface {
	Create(ctx context.Context, m *models.SubMenu) error
	Update(ctx context.Context, m *models.SubMenu) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.SubMenu])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.SubMenu])
}

type Repository interface {
	base.Repository
	Commond
	Query
}
