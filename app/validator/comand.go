package validator

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CmdForm struct {
	Cmd     string `json:"cmd"`
	Value   string `json:"value"`
	Options string `json:"options"`
}

func ProcessCmdForm(c *fiber.Ctx) CmdForm {
	var cmdForm CmdForm

	cmdForm.Cmd = c.FormValue("cmd")
	cmdForm.Options = c.FormValue("options")
	cmdForm.Value = c.FormValue("value")

	fmt.Println("cmd: " + cmdForm.Cmd + " options: " + cmdForm.Options + " value: " + cmdForm.Value)
	return cmdForm
}

func (f *CmdForm) CheckForReqFields() (bool, error) {
	switch f.Cmd {
	case "say":
		return hasValue(f.Cmd, f.Value)
	case "time":
		return hasValue(f.Cmd, f.Value)
	case "weather":
		validOptions := []string{"clear", "rain", "thunder"}
		return hasOption(f.Cmd, f.Options, validOptions)
	case "kick":
		return hasValue(f.Cmd, f.Value)
	case "op":
		return hasValue(f.Cmd, f.Value)
	case "deop":
		return hasValue(f.Cmd, f.Value)
	case "whitelist":
		validOptions := []string{"add", "remove"}
		value, err := hasValue(f.Cmd, f.Value)
		if err != nil {
			return false, err
		}
		if value {
			return hasOption(f.Cmd, f.Options, validOptions)
		}
	case "setworldspawn":
		return hasValue(f.Cmd, f.Value)
	case "tpspawn":
		return hasValue(f.Cmd, f.Value)
	default:
		return false, errors.New("no command found")
	}
	return false, errors.New("no command found in switch")
}
