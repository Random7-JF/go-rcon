package model

import "fmt"

func GetAllUsers() ([]Users, error) {
	var users []Users
	result := dbSession.Db.Find(&users)
	if result.Error != nil {
		return users, nil
	}
	return users, nil
}

func GetUserById(id int) (Users, error) {
	var user Users
	result := dbSession.Db.Find(&user, id)
	if result.Error != nil {
		return user, nil
	}
	fmt.Println("User: ", user)
	return user, nil
}

func GetUserByUsername(username string) (Users, error) {
	var user Users
	result := dbSession.Db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return user, nil
	}
	fmt.Println("User: ", user)
	return user, nil
}

func UpdateUser(id int) error {

	return nil
}
