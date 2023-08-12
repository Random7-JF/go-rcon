package rcon

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
)

// GetPlayers sends a command over the rcon connection and takes the response, parse the string and return
// the Current player count, Max player count and list of currently connected players in models.Players
func GetPlayers(App *config.App) (model.PlayersCommand, error) {
	var playersJson model.PlayersCommand

	if !App.Rcon.Connection {
		fmt.Println("RCON disconnected")
		return model.PlayersCommand{}, errors.New("rcon Disconnected")
	}
	cmdresp, err := App.Rcon.Session.SendCommand("list")

	if err != nil {
		log.Println("GetPlayers send command failed: ", err)
		RconSession.Connected = false
		return playersJson, err
	}

	parseStr := strings.Split(cmdresp, ":")

	if strings.Contains(parseStr[0], "max") { // for mc <= 1.18
		playersJson, err = ParseListNew(cmdresp)
	} else { // for mc > 1.18
		playersJson, err = ParseListOld(cmdresp)
	}

	if err != nil {
		log.Println("GetPlayers Parse Failed Failed: ", err)
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

// GetWhitelist sends the whitelist list command which returns a string of the current count of whitelisted players and the play names.
// the function the parses the string to pull the count out and convert it to an int, and populates the model.whitelistcommand.players with names of the players.
func GetWhitelist(App *config.App) (model.WhitelistCommand, error) {
	var whitelist model.WhitelistCommand
	if !App.Rcon.Connection {
		fmt.Println("RCON disconnected")
		return model.WhitelistCommand{}, errors.New("rcon Disconnected")
	}
	cmdresp, err := App.Rcon.Session.SendCommand("whitelist list")

	if err != nil {
		fmt.Println("SendCommand failed:", err)
		return whitelist, err
	}

	parseStr := strings.Split(cmdresp, ":")
	count := ParseForCount(parseStr[0])
	whitelist.Count, err = strconv.Atoi(count)
	if err != nil {
		fmt.Println("MaxCount AtoI Failed:", err)
		return whitelist, err
	}
	players := strings.Split(parseStr[1], ",")

	for _, s := range players {
		x := strings.TrimSuffix(s, "\n")
		whitelist.Players = append(whitelist.Players, model.Players{Name: strings.Trim(x, " ")})
	}

	return whitelist, nil
}

// KickPlayer send them kick command over the rcon session, the target is the players name who you wish to kick
// function returns a model.kickcommand and error. if there is an error it is inputted into model.kickcommand.Error
func KickPlayer(App *config.App, target string) (model.KickCommand, error) {
	var kickCommand model.KickCommand
	var err error

	cmd := fmt.Sprintf("kick " + target)
	kickCommand.Response, err = App.Rcon.Session.SendCommand(cmd)
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
func SendMessage(App *config.App, message string) (model.NoReplyCommand, error) {
	var response model.NoReplyCommand
	var err error

	msg := "say [Go-Rcon]: " + message
	msg = strings.Replace(msg, "%20", " ", -1)

	response.Error, err = App.Rcon.Session.SendCommand(msg)
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
func SetTime(App *config.App, time string) (model.CommandResponse, error) {
	var response model.CommandResponse
	var err error

	cmd := "time set " + time
	response.Response, err = App.Rcon.Session.SendCommand(cmd)
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
func StopServer(App *config.App, confirm bool) (model.NoReplyCommand, error) {
	var response model.NoReplyCommand

	if confirm {
		response.Error, _ = App.Rcon.Session.SendCommand("stop")
		go model.AddToCommandLog(model.CommandLog{
			CommandType: "stop",
			Command:     "stop",
			SentBy:      "api",
		})
	}

	return response, nil
}

// SetWeather sends the weather XXXX command to the server of the rcon connection
// this will switch the current weather on the server and then log it to the command log database.
func SetWeather(App *config.App, weather string) model.CommandResponse {
	var response model.CommandResponse
	var err error

	cmd := "weather " + weather
	response.Response, err = App.Rcon.Session.SendCommand(cmd)
	if err != nil {
		response.Error = err.Error()
		return response
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "weather",
		Command:     cmd,
		Response:    response.Response,
		SentBy:      "api",
	})

	return response
}

func SetWorldSpawn(App *config.App, coord string) model.CommandResponse {
	var resp model.CommandResponse
	var err error

	cmd := "setworldspawn " + coord
	resp.Response, err = App.Rcon.Session.SendCommand(cmd)
	if err != nil {
		resp.Error = err.Error()
		return resp
	}

	model.SetSpawnCoords(coord)
	if err != nil {
		resp.Error = err.Error()
		return resp
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "setworldspwan",
		Command:     cmd,
		Response:    resp.Response,
		SentBy:      "api",
	})

	return resp
}

func TpToSpawn(App *config.App, player string) (model.CommandResponse, error) {
	var resp model.CommandResponse

	server, err := model.GetServerSettings()
	if err != nil {
		resp.Error = err.Error()
		return resp, err
	}

	cmd := "tp " + player + " " + server.SpawnCoords
	resp.Response, err = App.Rcon.Session.SendCommand(cmd)
	if err != nil {
		resp.Error = err.Error()
		return resp, err
	}

	go model.AddToCommandLog(model.CommandLog{
		CommandType: "tpspawn",
		Command:     cmd,
		Response:    resp.Response,
		SentBy:      "api",
	})

	return resp, nil

}
