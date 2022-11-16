package errors

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ERR_SOMETHING_WRONG = "oops something wrong. dont worry we will fix it"
	ERR_NOT_FOUND       = "not found"
	ERR_DUPLICATE_KEY   = "duplicate key"
	ERR_ALREADY_EXISTS  = "is already exists"
	ERR_INVALID_KEY     = "invalid"
	ERR_NOT_AUTHORIZED  = "unauthorized"
)

func CustomError(err string) error {
	return errors.New(err)
}

func ErrRecordNotFound(logger *zap.Logger, entity string, err error) error {
	if err == nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprint(entity, " ", ERR_NOT_FOUND))
	}
	return ErrSomethingWrong(logger, err)
}

func ErrDuplicateValue(logger *zap.Logger, entity string, err error) error {
	if err == nil {
		return err
	}
	if strings.Contains(strings.ToLower(err.Error()), ERR_DUPLICATE_KEY) {
		return errors.New(fmt.Sprint(entity, " ", ERR_ALREADY_EXISTS))
	}
	return ErrSomethingWrong(logger, err)
}

func ErrSomethingWrong(logger *zap.Logger, err error) error {
	if err == nil {
		return err
	}
	logger.Error(err.Error())
	return errors.New(ERR_SOMETHING_WRONG)
}

func ErrInvalidAttributes(attributes string) error {
	return errors.New(fmt.Sprint(ERR_INVALID_KEY, " ", attributes))
}

func ErrUnAuthorized() error {
	return errors.New(ERR_NOT_AUTHORIZED)
}
