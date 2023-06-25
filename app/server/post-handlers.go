package server

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/validator"
	"github.com/gofiber/fiber/v2"
)

func PostCommandsHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/commands")
	}

	return c.Redirect("/commands")
}

func PostPlayersHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/players")
	}

	return c.Redirect("/players")
}

func PostWhitelistHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/whitelist")
	}

	return c.Redirect("/whitelist")

}

func PostLoginHandler(c *fiber.Ctx) error {
	session, err := AppConfig.Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	userForm := validator.ProcessUserForm(c)
	err = userForm.CheckForBlanks()

	if err != nil {
		session.Set("Auth", model.Auth{
			Status:  false,
			Message: "Enter Username / Password",
		})
		if err := session.Save(); err != nil {
			return err
		}
		return c.Redirect("/user/login")
	}

	err = model.Authenticate(userForm.User, userForm.Password)
	if err != nil {
		session.Set("Auth", model.Auth{
			Status:  false,
			Message: "Incorrect Username / Password",
		})
		if err := session.Save(); err != nil {
			return err
		}
		return c.Redirect("/user/login")
	}

	session.Set("Auth", model.Auth{
		Status:  true,
		Message: "Successful Login",
	})
	if err := session.Save(); err != nil {
		return err
	}
	return c.Redirect("/user/login")
}
