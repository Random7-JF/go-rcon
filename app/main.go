package main

import (
	"encoding/gob"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/Random7-JF/go-rcon/app/server"
)

var App config.App

func main() {
	registerGlobs()
	//Load app configuration from either the config.json or env variables
	App.SetupAppConfig()
	//Start the main functions of the app, connecting to the DB, connecting to the Rcon server
	go setupDB()
	go setupRcon()
	//and start the webserver to take requests
	server.Serve(&App)
}

// setupDB runs the SetupDB function, this updates our App variable and registers needed info inside it for
// reference in other functions. This function will panic if no database can be connected to.
func setupDB() {
	model.SetupDB(&App)
	dbsession := model.SetupDbSession(&App)
	model.NewDbSession(dbsession)
}

// setupRcon runs the SetupConnection function, this updates our App variable and registers need info inside
// it for reference in other functions. The program will run with this unable to connect but won't be
// able to use rcon functions.
func setupRcon() {
	rcon.ConnectSession(&App)
	//rcon.SetupConnection(&App)
	//rconsession := rcon.SetupRconSession(&App)
	//rcon.NewRconSession(rconsession)
}

func registerGlobs() {
	gob.Register(model.Auth{})
}
