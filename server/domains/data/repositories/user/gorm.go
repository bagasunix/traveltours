package user

import (
	"context"

	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/gofrs/uuid"
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
	return "User"
}

// GetByKeyword implements Repository
func (g *gormProvider) GetByKeywordEmail(ctx context.Context, entity string, limit int64) (result models.SliceResult[models.User]) {
	entities := "%" + entity + "%"
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Find(&result.Value, "email=?", entities).Limit(int(limit)).Error)
	return result
}

// UpdateByStatus implements Repository
func (g *gormProvider) UpdateByStatus(ctx context.Context, m *models.User) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Updates(m).Error)
}

// Create implements Repository
func (g *gormProvider) Create(ctx context.Context, m *models.User) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(m).Error)
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(models.NewUserBuilder().Build(), "id=?", id.String()).Error)
}

// Update implements Repository
func (g *gormProvider) Update(ctx context.Context, m *models.User) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Updates(m).Error)
}

// GetByAll implements Repository
func (g *gormProvider) GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.User]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Find(&result.Value).Limit(int(limit)).Error)
	return result
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.User]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id=?", id.String()).Error)
	return result
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logger = logger
	g.db = db
	return g
}
