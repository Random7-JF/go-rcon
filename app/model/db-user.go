package model

func GetAllUsers() ([]Users, error) {
	var users []Users
	result := dbSession.Db.Find(&users)
	if result.Error != nil {
		return users, nil
	}
	return users, nil
}
