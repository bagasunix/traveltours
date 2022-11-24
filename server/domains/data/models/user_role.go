package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserRoles struct {
	UserId    uuid.UUID `gorm:"primaryKey"`
	RoleId    uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
}
