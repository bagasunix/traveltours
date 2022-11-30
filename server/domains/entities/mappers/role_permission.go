package mappers

import (
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/gofrs/uuid"
)

func RolePermissionModelToEntity(model *models.RolePermission) *entities.RolePermission {
	if model == nil {
		return nil
	}

	builder := entities.NewRolePermissionBuilder()
	builder.SetRoleId(model.RoleId)
	builder.SetPermissionId(model.PermissionId)
	builder.SetCreatedBy(model.CreatedBy)
	builder.SetCreatedAt(model.CreatedAt)
	return builder.Build()
}

func ListRolePermissionModelToListEntity(models []models.RolePermission) (entities []entities.RolePermission) {
	for _, m := range models {
		entities = append(entities, *RolePermissionModelToEntity(&m))
	}
	return entities
}

func PermissionRoleIdFromRolePermission(user *entities.RolePermission) (id any) {
	return user.PermissionId
}

func ListPermissionRoleIdFromRolePermission(roles []entities.RolePermission) (ids []uuid.UUID) {
	for _, i := range roles {
		ids = append(ids, PermissionRoleIdFromRolePermission(&i).(uuid.UUID))
	}
	return ids
}
