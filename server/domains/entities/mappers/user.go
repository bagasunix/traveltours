package mappers

import (
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"github.com/bagasunix/traveltours/server/domains/entities"
	"github.com/gofrs/uuid"
)

func UserModelToEntity(model *models.User) *entities.User {
	if model == nil {
		return nil
	}
	builder := entities.NewUserBuilder()
	builder.SetId(model.Id)
	builder.SetEmail(model.Email)
	// builder.SetRole(model.Role.Name)
	// builder.SetRole(model.RoleId)
	// builder.SetStatus(model.Status)
	// builder.SetStatus(model.Status.Name)
	return builder.Build()
}

func ListUserModelToListEntity(models []models.User) (entities []entities.User) {
	for _, m := range models {
		entities = append(entities, *UserModelToEntity(&m))
	}
	return entities
}

func UserIdFromRole(user *entities.User) (id any) {
	return user.Id
}

func ListUserIdFromListRole(users []entities.User) (ids []uuid.UUID) {
	for _, i := range users {
		ids = append(ids, UserIdFromRole(&i).(uuid.UUID))
	}
	return ids
}
