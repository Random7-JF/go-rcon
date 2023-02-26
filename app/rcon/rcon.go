package rcon

import (
	"fmt"
	"strings"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"

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

// GetPlayers sends a command over the rcon connection and takes the response, parse the string and return
// the Current player count, Max player count and list of currently connected players in models.Players
func GetPlayers() (model.PlayersCommand, error) {
	var playersJson model.PlayersCommand
	cmdresp, err := RconSession.Rcon.SendCommand("list")

	if err != nil {
		fmt.Println("SendCommand Failed:", err)
		return playersJson, err
	}

	parseStr := strings.Split(cmdresp, ":")

	if strings.Contains(parseStr[0], "max") { // for mc <= 1.18
		playersJson, err = ParseListNew(cmdresp)
	} else { // for mc > 1.18
		playersJson, err = ParseListOld(cmdresp)
	}

	if err != nil {
		fmt.Println("Parse Failed Failed:", err)
		return playersJson, err
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "list",
		Command:     "list",
		Response:    cmdresp,
		SentBy:      "Api",
	})
	return playersJson, nil
}

// KickPlayer send them kick command over the rcon session, the target is the players name who you wish to kick
// function returns a model.kickcommand and error. if there is an error it is inputted into model.kickcommand.Error
func KickPlayer(target string) (model.KickCommand, error) {
	var kickCommand model.KickCommand
	var err error

	cmd := fmt.Sprintf("kick " + target)
	kickCommand.Response, err = RconSession.Rcon.SendCommand(cmd)
	if err != nil {
		kickCommand.Error = err.Error()
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "kick",
		Command:     cmd,
		Response:    kickCommand.Response,
		SentBy:      "Api",
	})

	return kickCommand, nil
}

// SendMessage send a message prefixed with "[Go-Rcon]" to the server for all players to see.
// Using strings.replace to replace any %20 with a space that come from the Params.
func SendMessage(message string) (model.NoReplyCommand, error) {
	var response model.NoReplyCommand
	var err error

	msg := "say [Go-Rcon]: " + message
	msg = strings.Replace(msg, "%20", " ", -1)

	response.Error, err = RconSession.Rcon.SendCommand(msg)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "say",
		Command:     msg,
		SentBy:      "api",
	})

	return response, nil
}

// SendMessage send a message prefixed with "[Go-Rcon]" to the server for all players to see.
// Using strings.replace to replace any %20 with a space that come from the Params.
func SetTime(time string) (model.CommandResponse, error) {
	var response model.CommandResponse
	var err error

	cmd := "time set " + time
	response.Response, err = RconSession.Rcon.SendCommand(cmd)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "time",
		Command:     cmd,
		SentBy:      "api",
		Response:    response.Response,
	})

	return response, nil
}

// StopServer send the "stop" command to the sever over the rcon connection
// this will tell the server to shutdown and it will shutdown. A managed server should start backup after
func StopServer(confirm bool) (model.NoReplyCommand, error) {
	var response model.NoReplyCommand

	if confirm {
		response.Error, _ = RconSession.Rcon.SendCommand("stop")
		go model.AddToCommandLog(model.CommandLog{
			CommandType: "stop",
			Command:     "stop",
			SentBy:      "api",
		})
	}

	return response, nil
}
