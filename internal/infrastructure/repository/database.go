package repository

import (
	"go_playground/internal/infrastructure/config"
	"go_playground/internal/infrastructure/repository/clickhouse"
	"go_playground/internal/infrastructure/repository/mysql"
	"go_playground/internal/infrastructure/repository/postgresql"

	"gorm.io/gorm"
)

type SQLDatabaseConnection interface {
	Connect(cfg *config.Config) *gorm.DB
}

func NewDBGetConnection(dbDriver string) SQLDatabaseConnection {
	switch dbDriver {
	case "mysql":
		return &mysql.MySQLConnection{}
	case "postgres":
		return &postgresql.PostgreSQLConnection{}
	case "clickhouse":
		return &clickhouse.ClickHouseConnection{}
	default:
		return nil
	}

}
