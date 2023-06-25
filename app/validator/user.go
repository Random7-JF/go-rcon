package validator

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type UserForm struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func ProcessUserForm(c *fiber.Ctx) UserForm {
	var userForm UserForm

	userForm.User = c.FormValue("username")
	userForm.Password = c.FormValue("password")

	return userForm
}

func (userForm *UserForm) CheckForBlanks() error {
	if len(userForm.User) == 0 || len(userForm.Password) == 0 {
		return errors.New("this form has a blank submission")
	}
	return nil
}
