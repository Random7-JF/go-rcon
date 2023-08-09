package rcon

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"

	mcrcon "github.com/Kelwing/mc-rcon"
)

type Connection struct {
	Rcon *mcrcon.MCConn
}

var RconSession *Connection

func SetupRconSession(a *config.App) *Connection {
	return &Connection{
		Rcon: a.Rcon,
	}
}

func NewRconSession(c *Connection) {
	RconSession = c
}

func SetupConnection(App *config.App) error {
	App.Rcon = new(mcrcon.MCConn)

	ip := App.RconSettings.Ip + ":" + App.RconSettings.Port
	err := App.Rcon.Open(ip, App.RconSettings.Password)
	if err != nil {
		fmt.Println("Error opening rcon connection:", err)
		return err
	}
	err = App.Rcon.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating rcon connection:", err)
		return err
	}

	test, err := App.Rcon.SendCommand("list")
	if err != nil {
		fmt.Println("Error sending command:", err)
		return err
	}

	fmt.Println(test)
	App.RconSettings.Connection = true
	return nil
}
