package model

import (
	"errors"
	"fmt"
	"log"

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

func DeleteUser(id uint) error {
	user, err := GetUserById(id)
	if err != nil {
		fmt.Println("Delete user error:", err)
	}

	result := dbSession.Db.Delete(&user)
	if result.Error != nil {
		fmt.Println("Delete user error:", result.Error)
	}

	return nil
}

func UpdateUserPass(id int, password string) error {
	curUser, _ := GetUserById(uint(id))
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	curUser.Password = string(hashedPassword)

	result := dbSession.Db.Save(&curUser)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Updated password for: ", curUser)
	return nil
}

func CreateUser(username string, password string) error {

	checkUser, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	if checkUser.UserName == username {
		log.Println("Username Already in use.")
		return errors.New("username already in use")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	newUser := Users{UserName: username, Password: string(hashedPassword), Admin: false}

	result := dbSession.Db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("User creation error:", result.Error)
	}

	fmt.Println("User created:", newUser)
	return nil
}

func CreateDefaultAdmin() error {
	err := CreateUser("admin", "admin")
	if err != nil {
		log.Printf("CreateDefaultAdmin - CreateUser: Unable to create defaut admin: %s \n", err.Error())
		return err
	}

	admin := SetUserAdmin("admin", true)
	if admin != nil {
		log.Printf("CreateDefaultAdmin - SetUserAdmin:  Unable assign admin: %s \n", err.Error())
		return err
	}

	log.Printf("---- Created Default Admin Account ----")
	log.Printf("---- Username: admin               ----")
	log.Printf("---- Password: admin               ----")

	return nil
}

func IsAdminAccountCreated() bool {
	var user Users
	result := dbSession.Db.Where("admin = ?", true).First(&user)
	log.Println(result.RowsAffected)
	log.Println(result.Error)

	return result.Error != nil
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

func IsUserAdmin(username string) bool {
	result, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	return result.Admin
}

func SetUserAdmin(username string, admin bool) error {
	result, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	result.Admin = admin
	query := dbSession.Db.Save(&result)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
