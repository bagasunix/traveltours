package mappers

import (
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/entities"
)

func RoleModelToEntity(model *models.Role) *entities.Role {
	if model == nil {
		return nil
	}
	builder := entities.NewRoleBuilder()
	builder.SetName(model.Name)
	builder.SetGroup(model.Group)
	builder.SetCreatedAt(model.CreatedAt)
	builder.SetCreatedBy(model.CreatedBy)
	return builder.Build()
}

func ListRoleModelToListEntity(models []models.Role) (entities []entities.Role) {
	for _, m := range models {
		entities = append(entities, *RoleModelToEntity(&m))
	}
	return entities
}
