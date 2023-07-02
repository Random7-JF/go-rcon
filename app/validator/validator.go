package validator

import "github.com/gofiber/fiber/v2"

type iForm interface {
	ProcessForm(f interface{})
	EmptyFields() bool
	ValidateFields() bool
}

type CmdForm struct {
	Cmd     string `json:"cmd"`
	Value   string `json:"value"`
	Options string `json:"options"`
}

type UserForm1 struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Action   string `json:"action"`
}

type AdminForm struct {
	User  string `json:"user"`
	Admin bool   `json:"admin:`
}

type LoginForm struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func Validate(f *iForm) interface{} {
	return f
}

func (f *LoginForm) ProcessForm(c *fiber.Ctx) {
	f.User = c.FormValue("username")
	f.Password = c.FormValue("password")
}

func (f *LoginForm) EmptyFields() bool {
	if len(f.User) != 0 || len(f.Password) != 0 {
		return true
	}
	return false
}

func (f *LoginForm) ValidateFields() bool {
	//check for minimum characters

	//check for bad characters

	return true
}
