package permission

import (
	"context"
	"fmt"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/pkg/helpers"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type gormProvider struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (g *gormProvider) GetBySelectedId(ctx context.Context, ids []uuid.UUID) (result models.SliceResult[models.Permission]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id in ?", helpers.ListUUIDToListString(ids)).Find(&result.Value).Error)
	return result
}

func (g *gormProvider) GetByKeywords(ctx context.Context, keyword string, limit int64) (result models.SliceResult[models.Permission]) {
	q := fmt.Sprint('%', keyword, '%')
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("name like ?", q).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(new(models.Permission), "id = ?", id.String()).Error)
}

func (g *gormProvider) GetModelName() string {
	return "permission"
}

func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Permission]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id = ?", id.String()).Error)
	return result
}

func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Permission]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

func (g *gormProvider) Create(ctx context.Context, model *models.Permission) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(model).Error)
}

func (g *gormProvider) Update(ctx context.Context, model *models.Permission) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Updates(model).Error)
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	a := new(gormProvider)
	a.logger = logger
	a.db = db
	return a
}
