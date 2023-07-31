package model

import (
	"github.com/Random7-JF/go-rcon/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Db *gorm.DB
}

var dbSession *Connection

func SetupDbSession(a *config.App) *Connection {
	return &Connection{
		Db: a.Db,
	}
}

func NewDbSession(c *Connection) {
	dbSession = c
}

func SetupDB(App *config.App) {
	var err error
	dsn := "host=" + App.DbSettings.Host + " user=" + App.DbSettings.User + " password=" + App.DbSettings.Password + " dbname=" + App.DbSettings.DbName + " port=" + App.DbSettings.Port + " sslmode=" + App.DbSettings.Sslmode + " TimeZone=America/Winnipeg"
	App.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	App.Db.AutoMigrate(&Users{})
	App.Db.AutoMigrate(&CommandLog{})

}
