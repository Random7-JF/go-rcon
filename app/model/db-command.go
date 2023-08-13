package model

import (
	"fmt"
	"log"
)

// AddToCommandLog takes a Commandlog and enters it into the database.
func AddToCommandLog(log CommandLog) error {

	result := dbSession.Db.Create(&log)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return result.Error
	}

	return nil
}

// GetCommandLog returns a []Commandlog containing the Limit of rows in the Commandlog table, excluding list commands, ordered descened.
// set limit to 0 to retrieve all commands
func GetCommandLog(limit int) ([]CommandLog, error) {
	var cmdlog []CommandLog
	if limit > 0 {
		query := dbSession.Db.Limit(limit).Where("command != ?", "list").Order("created_at DESC").Find(&cmdlog)
		if query.Error != nil {
			log.Println(query.Error)
			return cmdlog, query.Error
		}
	} else {
		query := dbSession.Db.Where("command != ?", "list").Order("created_at DESC").Find(&cmdlog)
		if query.Error != nil {
			log.Println(query.Error)
			return cmdlog, query.Error
		}
	}
	return cmdlog, nil
}

func SetSpawnCoords(coords string) error {
	var serverSetting ServerSettings

	result := dbSession.Db.Last(&serverSetting)
	if result.Error != nil {
		result := dbSession.Db.Create(&serverSetting)
		if result.Error != nil {
			log.Println(result.Error.Error())
			return result.Error
		}
		return nil
	}

	serverSetting.SpawnCoords = coords
	dbSession.Db.Save(&serverSetting)

	return nil
}

func SetRconSettings(ip string, port string, password string) error {
	var serverSettings ServerSettings

	result := dbSession.Db.Last(&serverSettings)
	if result.Error != nil {
		result := dbSession.Db.Create(&serverSettings)
		if result.Error != nil {
			log.Println(result.Error.Error())
			return result.Error
		}
		return nil
	}

	serverSettings.RconIp = ip
	serverSettings.RconPort = port
	serverSettings.RconPass = password
	dbSession.Db.Save(&serverSettings)
	return nil
}

func GetServerSettings() (ServerSettings, error) {
	var serverSettings ServerSettings

	query := dbSession.Db.Last(&serverSettings)
	if query.Error != nil {
		log.Println(query.Error)
		return serverSettings, query.Error
	}
	return serverSettings, nil
}
