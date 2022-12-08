package userstatus

import (
	"context"

	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/base"
)

type Query interface {
	GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.UserStatus])
	GetById(ctx context.Context, id int8) (result models.SingleResult[*models.UserStatus])
}

type Repository interface {
	base.Repository
	Query
}
