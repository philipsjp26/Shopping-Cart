package clickhouse

import (
	"fmt"
	"go_playground/internal/infrastructure/config"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/clickhouse"

	"gorm.io/gorm"
)

type ClickHouseConnection struct{}

func (c *ClickHouseConnection) Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("tcp://%s:%s?database=%s&username=%spassword=%s&read_timeout=%v&write_timeout=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Username,
		cfg.Database.Password,
		strconv.Itoa(cfg.Database.ReadTimeout),
		strconv.Itoa(cfg.Database.WriteTimeout),
	)
	db, err := gorm.Open(clickhouse.Open(dsn))
	if err != nil {
		log.Error(err)
	}
	return db

}
