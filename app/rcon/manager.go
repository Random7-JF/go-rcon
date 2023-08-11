package rcon

import (
	"fmt"
	"log"

	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/Random7-JF/go-rcon/app/config"
)

func ConnectSession(App *config.App) error {
	App.Rcon.Session = new(mcrcon.MCConn)

	ip := App.Rcon.Ip + ":" + App.Rcon.Port
	err := App.Rcon.Session.Open(ip, App.Rcon.Password)
	if err != nil {
		fmt.Println("Error opening rcon connection:", err)
		return err
	}
	err = App.Rcon.Session.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating rcon connection:", err)
		return err
	}

	TestSession(App)
	App.Rcon.Connection = true
	return nil
}

func DisconnectSession(App *config.App) {
	App.Rcon.Session.Close()
}

func TestSession(App *config.App) {
	test, err := App.Rcon.Session.SendCommand("list")
	if err != nil {
		log.Println("Error sending command:", err)
		App.Rcon.Connection = false
	}

	log.Println("Test Session - Session: " + test)
	App.Rcon.Connection = true
}
