package postgres

import (
	"fmt"
	"log"

	"github.com/saufiroja/order-service/config"
	"github.com/saufiroja/order-service/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", conf.Postgres.DB_HOST, conf.Postgres.DB_USER, conf.Postgres.DB_PASS, conf.Postgres.DB_NAME, conf.Postgres.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database")
	}

	log.Println("Connected to database")

	db.AutoMigrate(&entity.Order{})

	return db
}
