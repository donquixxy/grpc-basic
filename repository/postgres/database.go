package postgres

import (
	"fmt"
	"gokit-basic/common/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(conf *config.AppConfig) *gorm.DB {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.DatabaseConf.Address, conf.DatabaseConf.Username, conf.Password, conf.DbName, conf.DatabaseConf.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}

	e, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	p := e.Ping()

	if p != nil {
		log.Fatalf("failed to ping db %v", p)
	}

	return db
}
