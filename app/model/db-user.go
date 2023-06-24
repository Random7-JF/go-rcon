package model

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]Users, error) {
	var users []Users
	result := dbSession.Db.Find(&users)
	if result.Error != nil {
		return users, nil
	}
	return users, nil
}

func GetUserById(id uint) (Users, error) {
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

func Authenticate(username, testPassword string) (uint, string, error) {

	result, err := GetUserByUsername(username)
	if err != nil {
		return 0, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}
	fmt.Println("Logged in")
	return result.ID, result.Password, nil
}
