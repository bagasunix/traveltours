package mappers

import (
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/entities"
)

func PermissionModelToEntity(model *models.Permission) *entities.Permission {
	if model == nil {
		return nil
	}
	builder := entities.NewPermissionBuilder()
	builder.SetName(model.Name)
	builder.SetMethod(model.Method)
	for _, v := range model.Url {
		builder.SetUrl(v)
	}
	builder.SetCreatedBy(model.CreatedBy)
	builder.SetCreatedAt(model.CreatedAt)
	return builder.Build()
}

func ListPermissionModelToListEntity(models []models.Permission) (entities []entities.Permission) {
	for _, m := range models {
		entities = append(entities, *PermissionModelToEntity(&m))
	}
	return entities
}
