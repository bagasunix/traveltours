package helpers

import (
	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

func GenerateUUIDV1(logger *zap.Logger) uuid.UUID {
	id, err := uuid.NewV1()
	errors.HandlerWithLoggerReturnedVoid(logger, err, "uuid", "generator")
	return id
}

func GenerateUUIDV4(logger *zap.Logger) uuid.UUID {
	id, err := uuid.NewV4()
	errors.HandlerWithLoggerReturnedVoid(logger, err, "uuid", "generator")
	return id
}
