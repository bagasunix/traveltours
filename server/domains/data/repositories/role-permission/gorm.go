package rolepermission

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

// CreateBatch implements Repository
func (g *gormProvider) CreateBatch(ctx context.Context, models []models.RolePermission) error {
	for _, v := range models {
		if err := g.Create(ctx, &v); err != nil {
			return err
		}
	}
	return nil
}

// CreateTx implements Repository
func (g *gormProvider) CreateTx(ctx context.Context, tx any, model *models.RolePermission) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), tx.(*gorm.DB).WithContext(ctx).Create(tx).Error)
}

// CreateTxBatch implements Repository
func (g *gormProvider) CreateTxBatch(ctx context.Context, tx any, models []models.RolePermission) error {
	for _, v := range models {
		if err := g.CreateTx(ctx, tx, &v); err != nil {
			return err
		}
	}
	return nil
}

// DeleteBatch implements Repository
func (g *gormProvider) DeleteBatch(ctx context.Context, roleId uuid.UUID, permissionIds []uuid.UUID) error {
	for _, v := range permissionIds {
		if err := g.Delete(ctx, roleId, v); err != nil {
			return err
		}
	}
	return nil
}

// GetById implements Repository
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SliceResult[models.RolePermission]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id=?", id.String()).Error)
	return result
}

// GetConnection implements Repository
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository
func (g *gormProvider) GetModelName() string {
	return "Role Permission"
}

// GetByRoleId implements Repository
func (g *gormProvider) GetByRoleId(ctx context.Context, roleId uuid.UUID) (result models.SliceResult[*models.RolePermission]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("role_id = ?", roleId.String()).Find(&result.Value).Error)
	return result
}

// GetBySelectedRoleId implements Repository
func (g *gormProvider) GetBySelectedRoleId(ctx context.Context, roleIds []uuid.UUID) (result models.SliceResult[models.RolePermission]) {
	rids := helpers.ListUUIDToListString(roleIds)
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id in ?", rids).Find(&result.Value).Error)
	return result
}

// Create implements Repository
func (g *gormProvider) Create(ctx context.Context, m *models.RolePermission) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(m).Error)
}

// Delete implements Repository
func (g *gormProvider) Delete(ctx context.Context, roleId uuid.UUID, permissionId uuid.UUID) error {
	return errors.ErrSomethingWrong(g.logger,
		g.db.WithContext(ctx).Delete(models.NewRolePermissionBuilder().Build(),
			"role_id = ? and permission_id = ?", roleId.String(), permissionId.String()).Error)
}

// Update implements Repository
func (g *gormProvider) Update(ctx context.Context, m *models.RolePermission) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Updates(m).Error)
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	a := new(gormProvider)
	a.logger = logger
	a.db = db
	return a
}
