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
	builder.SetId(model.Id)
	builder.SetName(model.Name)
	builder.SetSlug(model.Slug)
	builder.SetMethod(model.Method)
	builder.SetUrl(model.Url)
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
