package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Random7-JF/go-rcon/app/rcon"
)

type CmdForm struct {
	Cmd     string `json:"cmd"`
	Value   string `json:"value"`
	Options string `json:"options"`
}

func (c *CmdForm) ValidateInputs() error {
	err := c.CheckForBlanks()
	if err != nil {
		return err
	}
	err = c.CheckCmd()
	if err != nil {
		return err
	}
	return nil
}

func (c *CmdForm) CheckForBlanks() error {
	if len(c.Cmd) != 0 && len(c.Value) != 0 {
		return nil
	}
	err := errors.New("this form has a blank submission")
	return err
}

func (c *CmdForm) CheckCmd() error {

	switch c.Cmd {
	case "say":
		_, err := rcon.SendMessage(c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "time":
		_, err := rcon.SetTime(c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "weather":
		_, err := rcon.RconSession.Rcon.SendCommand("weather " + c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "whitelist":
		_, err := rcon.RconSession.Rcon.SendCommand("whitelist add " + c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "kick":
		_, err := rcon.RconSession.Rcon.SendCommand("kick " + c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	case "op":
		resp, err := rcon.RconSession.Rcon.SendCommand("op " + c.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.Contains(resp, "already is an operator") {
			_, err = rcon.RconSession.Rcon.SendCommand("deop " + c.Value)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	default:
		err := errors.New("no command found")
		return err
	}
}
