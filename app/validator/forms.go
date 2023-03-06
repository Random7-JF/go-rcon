package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/gofiber/fiber/v2"
)

type Form struct {
	Cmd     string `json:"cmd"`
	Value   string `json:"value"`
	Options string `json:"options"`
}

func ProcessForm(c *fiber.Ctx) Form {
	var form Form
	form.Cmd = c.FormValue("cmd")
	form.Value = c.FormValue("value")
	form.Options = c.FormValue("options")

	return form
}

func (form *Form) ValidateInputs() error {
	err := form.CheckForBlanks()
	if err != nil {
		return err
	}
	err = form.CheckCmd()
	if err != nil {
		return err
	}
	return nil
}

func (form *Form) CheckForBlanks() error {
	if len(form.Cmd) != 0 && len(form.Value) != 0 {
		return nil
	}
	err := errors.New("this form has a blank submission")
	return err
}

func (form *Form) CheckCmd() error {

	switch form.Cmd {
	case "say":
		_, err := rcon.SendMessage(form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "time":
		_, err := rcon.SetTime(form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "weather":
		_, err := rcon.RconSession.Rcon.SendCommand("weather " + form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "whitelist":
		_, err := rcon.RconSession.Rcon.SendCommand("whitelist add " + form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "kick":
		_, err := rcon.RconSession.Rcon.SendCommand("kick " + form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "op":
		resp, err := rcon.RconSession.Rcon.SendCommand("op " + form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.Contains(resp, "already is an operator") {
			_, err = rcon.RconSession.Rcon.SendCommand("deop " + form.Value)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	case "dewhitelist":
		resp, err := rcon.RconSession.Rcon.SendCommand("whitelist remove " + form.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("whitelist remove:" + resp)
		return nil
	default:
		err := errors.New("no command found")
		return err
	}
}
