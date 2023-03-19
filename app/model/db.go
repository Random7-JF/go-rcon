package model

import (
	"fmt"

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
	dsn := "host=" + App.DbSettings.Host + " user=" + App.DbSettings.User + " password=" + App.DbSettings.Password + " dbname=" + App.DbSettings.DbName + " port=" + App.DbSettings.Port + " sslmode=disable TimeZone=America/Winnipeg"
	App.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	App.Db.AutoMigrate(&Users{})
	App.Db.AutoMigrate(&CommandLog{})

}

func AddToCommandLog(log CommandLog) error {

	result := dbSession.Db.Create(&log)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

func GetCommandLog(limit int) ([]CommandLog, error) {
	var cmdlog []CommandLog
	if limit > 0 {
		query := dbSession.Db.Limit(limit).Where("command != ?", "list").Order("created_at DESC").Find(&cmdlog)
		if query.Error != nil {
			fmt.Println(query.Error)
			return cmdlog, query.Error
		}
	} else {
		query := dbSession.Db.Where("command != ?", "list").Order("created_at DESC").Find(&cmdlog)
		if query.Error != nil {
			fmt.Println(query.Error)
			return cmdlog, query.Error
		}
	}
	fmt.Println(cmdlog)
	return cmdlog, nil
}
