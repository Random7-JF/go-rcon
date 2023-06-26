package server

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/validator"
	"github.com/gofiber/fiber/v2"
)

func PostCommandsHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/commands")
	}

	return c.Redirect("/app/commands")
}

func PostPlayersHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/players")
	}

	return c.Redirect("/app/players")
}

func PostWhitelistHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/whitelist")
	}

	return c.Redirect("/app/whitelist")

}

func PostLoginHandler(c *fiber.Ctx) error {
	userForm := validator.ProcessUserForm(c)
	err := userForm.CheckForBlanks()

	if err != nil {
		helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
			Status:  false,
			Message: "Enter Username / Password",
		})
		return c.Redirect("/login")
	}

	err = model.Authenticate(userForm.User, userForm.Password)
	if err != nil {
		helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
			Status:  false,
			Message: "Incorrect Username / Password",
		})
		return c.Redirect("/login")
	}

	helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
		Status:  true,
		Message: "Successful Login",
		Admin:   model.IsUserAdmin(userForm.User),
	})

	return c.Redirect("/app/dashboard")
}
