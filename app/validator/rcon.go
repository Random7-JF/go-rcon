package validator

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RconFrom struct {
	Ip       string
	Port     string
	Password string
}

type RconSession struct {
	Action string
}

func ProcessRconForm(c *fiber.Ctx) RconFrom {
	var rconForm RconFrom

	rconForm.Ip = c.FormValue("ip")
	rconForm.Port = c.FormValue("port")
	rconForm.Password = c.FormValue("pass")

	log.Println("ip: " + rconForm.Ip + " port: " + rconForm.Port + " password: " + rconForm.Password)
	return rconForm
}

func (r *RconFrom) CheckForReqFields() error {
	if !isBlank(r.Ip) || !isBlank(r.Port) || !isBlank(r.Password) {
		log.Printf("IP: %d Port: %d Password: %d", len(r.Ip), len(r.Port), len(r.Password))
		return errors.New("blank submission")
	} else {
		return nil
	}
}

func ProcessRconSessionForm(c *fiber.Ctx) RconSession {
	var rconSessionForm RconSession

	rconSessionForm.Action = c.FormValue("action")
	return rconSessionForm
}

func (r *RconSession) CheckForReqFields() error {
	valid, err := hasValue("action", r.Action)
	if err != nil {
		log.Println("CheckForReqFields - hasValue errored: ", err)
	}
	if valid {
		valid, err := hasOption("action", r.Action, []string{"stop", "start", "restart"})
		if err != nil {
			log.Println("CheckForReqFields - hasOption errored: ", err)
		}
		if valid {
			return nil
		}
	}
	return err
}
