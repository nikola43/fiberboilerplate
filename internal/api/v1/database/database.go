package database

import (
	"fmt"
	"log"

	config "github.com/nikola43/fiberboilerplate/internal/api/v1/config"
	dbmodels "github.com/nikola43/fiberboilerplate/internal/api/v1/models/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&dbmodels.User{})

	// CREATE
	GormDB.AutoMigrate(&dbmodels.User{})
}

func InitializeDatabase(databaseConfig config.DatabaseConfig) {
	var err error
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Name,
	)

	GormDB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}
}
