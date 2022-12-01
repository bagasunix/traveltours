package mappers

import (
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/gofrs/uuid"
)

func RoleModelToEntity(model *models.Role) *entities.Role {
	if model == nil {
		return nil
	}
	builder := entities.NewRoleBuilder()
	builder.SetId(model.Id)
	builder.SetName(model.Name)
	builder.SetGroup(model.Group)
	builder.SetCreatedAt(model.CreatedAt)
	builder.SetCreatedBy(model.CreatedBy)
	builder.SetIsActive(model.IsActive)
	return builder.Build()
}

func ListRoleModelToListEntity(models []models.Role) (entities []entities.Role) {
	for _, m := range models {
		entities = append(entities, *RoleModelToEntity(&m))
	}
	return entities
}

func RoleIdFromRole(user *entities.Role) (id any) {
	return user.Id
}

func ListRoleIdFromListRole(roles []entities.Role) (ids []uuid.UUID) {
	for _, i := range roles {
		ids = append(ids, RoleIdFromRole(&i).(uuid.UUID))
	}
	return ids
}
