package model

import "fmt"

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
	return cmdlog, nil
}
