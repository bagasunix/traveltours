package db

import (
	"context"
	"time"

	"github.com/bagasunix/traveltours/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewDB(ctx context.Context, logger *zap.Logger, dbConfig *DbConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	errors.HandlerWithOSExit(logger, err, "init", "database", "config", dbConfig.GetDSN())
	errors.HandlerWithOSExit(logger, db.WithContext(ctx).Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(dbConfig.MaxIdleConns).SetMaxOpenConns(dbConfig.MaxOpenConns).SetConnMaxLifetime(time.Hour)), "db_resolver")
	return db
}
