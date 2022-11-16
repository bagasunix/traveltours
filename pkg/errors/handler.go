package errors

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func HandlerWithOSExit(logger *zap.Logger, err error, args ...interface{}) {
	if err == nil {
		return
	}
	logger.Log(zap.FatalLevel, err.Error(), zap.Any("args", err))
	os.Exit(1)
}

func HandlerWithLoggerReturnedError(logger *zap.Logger, err error, args ...interface{}) error {
	if err == nil {
		return err
	}
	logger.Log(zap.ErrorLevel, err.Error(), zap.Any("arg", args))
	return err
}

func HandlerWithLoggerReturnedVoid(logger *zap.Logger, err error, args ...interface{}) {
	if err == nil {
		return
	}
	logger.Log(zap.ErrorLevel, err.Error(), zap.Any("arg", args))
}

func HandlerReturnedVoid(err error, args ...interface{}) {
	if err == nil {
		return
	}
	fmt.Println("err", err)
}
