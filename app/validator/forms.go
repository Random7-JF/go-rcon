package validator

import (
	"errors"
	"fmt"

	"github.com/Random7-JF/go-rcon/app/rcon"
)

type CmdForm struct {
	Cmd    string `json:"cmd"`
	Value  string `json:"value"`
	Params string `json:"params"`
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
	if len(c.Cmd) != 0 && len(c.Params) != 0 {
		return nil
	}

	fmt.Println(c.Cmd)
	fmt.Println(c.Params)

	err := errors.New("this form has a blank submission")
	return err
}

func (c *CmdForm) CheckCmd() error {
	switch c.Cmd {
	case "say":
		fmt.Println("Say Command Found")
		rcon.SendMessage(c.Params)
		return nil
	case "time":
		fmt.Println("Time Command Found")
		resp, err := rcon.SetTime(c.Params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		return nil
	case "weather":
		fmt.Println("Weather Command Found")
		fmt.Println(c.Params)
		resp, err := rcon.RconSession.Rcon.SendCommand("weather " + c.Params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		return nil
	default:
		err := errors.New("no command found")
		return err
	}
}
