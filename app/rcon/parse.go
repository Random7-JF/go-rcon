package rcon

import (
	"log"
	"strconv"
	"strings"

	"github.com/Random7-JF/go-rcon/app/model"
)

// ParseListOld is for the /list command pre-1.18 mc
func ParseListOld(cmdresp string) (model.PlayersCommand, error) {
	var playersJson model.PlayersCommand
	var err error
	index := strings.Index(cmdresp, "/")   // find "/" index
	countstr := cmdresp[index-2 : index+3] // substring based off index
	count := strings.Split(countstr, "/")  // split on "/"

	playerslist := strings.Split(cmdresp, ":")    // split at colon "There are 2/20 players online:Random777, Dude1872"
	players := strings.Split(playerslist[1], ",") //split at comma "Random777, Dude1872"

	playersJson.Players = ParseForPlayers(players)

	playersJson.CurrentCount, err = strconv.Atoi(strings.Trim(count[0], " "))
	if err != nil {
		log.Printf("ParseListOld - AtoI current count: Error: %v", err)
		return playersJson, err

	}

	playersJson.MaxCount, err = strconv.Atoi(strings.Trim(count[1], " "))
	if err != nil {
		log.Printf("ParseListOld - AtoI max count: Error: %v", err)
		return playersJson, err
	}

	return playersJson, nil
}

// Parselist/new is for the /list command <= 1.18 mc
func ParseListNew(cmdresp string) (model.PlayersCommand, error) {
	var playersJson model.PlayersCommand
	var err error
	parseStr := strings.Split(cmdresp, ":")
	countStr := strings.Split(parseStr[0], "max")

	playersJson.CurrentCount, err = strconv.Atoi(ParseForCount(countStr[0]))
	if err != nil {
		log.Printf("ParseListNew - AtoI current count: Error: %v", err)
		return playersJson, err

	}
	playersJson.MaxCount, err = strconv.Atoi(ParseForCount(countStr[1]))
	if err != nil {
		log.Printf("ParseListNew - AtoI max count: Error: %v", err)
		return playersJson, err
	}

	players := strings.Split(parseStr[1], ",") //split at comma "Random777, Dude1872"
	playersJson.Players = ParseForPlayers(players)

	return playersJson, nil
}

// ParsePlayers takes a string slice of all the players and splits, trims spaces and newlines and returns a slice of model.Players
func ParseForPlayers(p []string) []model.Players {
	var Players []model.Players
	for _, s := range p {
		player := strings.TrimSuffix(s, "\n")
		Players = append(Players, model.Players{Name: strings.Trim(player, " ")})
	}
	return Players
}

// ParseForCount takes a string containing the counts for current or max players and returns the number in a string
func ParseForCount(countString string) string {
	var result string
	for _, s := range countString {
		if strings.ContainsAny(string(s), "0123456789") {
			result = result + string(s)
		}
	}
	return result
}
