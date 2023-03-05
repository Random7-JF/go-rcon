package server

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/validator"
	"github.com/gofiber/fiber/v2"
)

func CmdHandler(c *fiber.Ctx) error {
	var Submitted validator.CmdForm

	Submitted.Cmd = c.FormValue("cmd")
	Submitted.Value = c.FormValue("value")
	Submitted.Options = c.FormValue("options")

	err := Submitted.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/commands")
	}

	return c.Redirect("/commands")
}

func PlayerCmdHandler(c *fiber.Ctx) error {
	var Submitted validator.CmdForm

	Submitted.Cmd = c.FormValue("cmd")
	Submitted.Value = c.FormValue("value")
	Submitted.Options = c.FormValue("options")

	err := Submitted.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/players")
	}

	return c.Redirect("/players")
}

func WhitelistCmdHandler(c *fiber.Ctx) error {
	var Submitted validator.CmdForm

	Submitted.Cmd = c.FormValue("cmd")
	Submitted.Value = c.FormValue("value")
	Submitted.Options = c.FormValue("options")

	err := Submitted.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/whitelist")
	}

	return c.Redirect("/whitelist")
}
