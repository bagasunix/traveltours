package userrole

import (
	"context"

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

// CreateTx implements Repository
func (g *gormProvider) CreateTx(ctx context.Context, tx any, m *models.Role) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), tx.(*gorm.DB).WithContext(ctx).Create(m).Error)
}

// GetByGroup implements Repository
func (g *gormProvider) GetByGroup(ctx context.Context, group string) (result models.SliceResult[models.Role]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("group = ?", group).Find(&result.Value).Error)
	return result
}

// GetBySelectedId implements Repository
func (g *gormProvider) GetBySelectedId(ctx context.Context, ids []uuid.UUID) (result models.SliceResult[models.Role]) {
	rids := helpers.ListUUIDToListString(ids)
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id in ?", rids).Find(&result.Value).Error)
	return result
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "Role"
}

// GetByKeywords implements Repository
func (g *gormProvider) GetByKeywords(ctx context.Context, entity string, limit int64) (result models.SliceResult[models.Role]) {
	entities := "%" + entity + "%"
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value, "name=?", entities).Error)
	return result
}

// Create implements Repository
func (g *gormProvider) Create(ctx context.Context, m *models.Role) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(m).Error)
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(models.NewRoleBuilder().Build(), "id=>", id.String()).Error)
}

// Update implements Repository
func (g *gormProvider) Update(ctx context.Context, m models.Role) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Updates(m).Error)
}

// GetByAll implements Repository
func (g *gormProvider) GetByAll(ctx context.Context, limit int64) (result models.SliceResult[models.Role]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Find(&result.Value).Error)
	return result
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Role]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id=?", id.String()).Error)
	return result
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logger = logger
	return g
}
