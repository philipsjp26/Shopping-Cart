package mysql

import (
	"fmt"
	"go_playground/internal/infrastructure/config"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnection struct {
}

func (m *MySQLConnection) Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Charset,
		cfg.Database.Timezone,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:        dsn,
		DriverName: "mysql",
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
	return db
}
