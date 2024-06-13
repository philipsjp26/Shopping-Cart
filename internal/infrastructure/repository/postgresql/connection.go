package postgresql

import (
	"fmt"
	"go_playground/internal/infrastructure/config"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type PostgreSQLConnection struct{}

func (p *PostgreSQLConnection) Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.Timezone,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err)

	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime))
	return nil
}
