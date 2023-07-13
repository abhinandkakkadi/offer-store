package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "github.com/abhinandkakkadi/offer-store/pkg/config"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	fmt.Printf("%+v", cfg)
	dsn := cfg.DB_USERNAME + ":" + cfg.DB_PASSWORD + "@tcp" + "(" + cfg.DB_HOST + ":" + cfg.DB_PORT + ")/" + cfg.DB_NAME + "?" + "parseTime=true&loc=Local"
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	fmt.Println(db, dbErr)

	return db, dbErr

}
