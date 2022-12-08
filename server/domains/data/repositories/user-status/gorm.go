package userstatus

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type gormProvider struct {
	logger *zap.Logger
	db     *gorm.DB
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "Status User"
}

// GetByAll implements Repository
func (g *gormProvider) GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.UserStatus]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Find(&result.Value).Limit(int(limit)).Error)
	return result
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id int8) (result models.SingleResult[*models.UserStatus]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id=?", id).Error)
	return result
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logger = logger
	g.db = db
	return g
}
