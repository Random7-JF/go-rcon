package rcon

import (
	"log"

	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
)

func ConnectSession(App *config.App) error {
	App.Rcon.Session = new(mcrcon.MCConn)

	rconSettings, err := model.GetServerSettings()
	if err != nil {
		log.Println("ConnectSession - GetSeverSettings: Error getting server settings: ", err)
	}

	ip := rconSettings.RconIp + ":" + rconSettings.RconPort
	err = App.Rcon.Session.Open(ip, rconSettings.RconPass)

	if err != nil {
		log.Println("ConnectSession - Open: Error opening rcon connection: ", err)
		return err
	}
	err = App.Rcon.Session.Authenticate()
	if err != nil {
		log.Println("ConnectSession - Authenticate: Error authenticating rcon connection: ", err)
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
		log.Println("TestSession - SendCommand: Error sending command: ", err)
		App.Rcon.Connection = false
	}

	log.Println("TestSession - SendCommand: " + test)
	App.Rcon.Connection = true
}
