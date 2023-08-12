package rcon

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"

	mcrcon "github.com/Kelwing/mc-rcon"
)

func SetupConnection(App *config.App) error {
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
