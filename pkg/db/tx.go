package db

import (
	"database/sql"

	"gorm.io/gorm"
)

type TxProvider interface {
	sql.Tx | *gorm.DB
}
