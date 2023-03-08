package server

import (
	"fmt"

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
