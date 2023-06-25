package model

import (
	"errors"

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
	return user, nil
}

func GetUserByUsername(username string) (Users, error) {
	var user Users
	result := dbSession.Db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return user, nil
	}
	return user, nil
}

func UpdateUser(id int) error {

	return nil
}

func Authenticate(username, testPassword string) error {
	result, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("incorrect password")
		return err
	} else if err != nil {
		return err
	}

	return nil
}
