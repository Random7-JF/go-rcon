package server

import (
	"github.com/Random7-JF/go-rcon/app/config"
)

func SetupRoutes(App *config.App) {
	//View routes
	//Get
	App.WebServer.Get("/", IndexHandler)
	App.WebServer.Get("/dashboard", DashboardHandler)
	App.WebServer.Get("/players", PlayersPageHandler)
	App.WebServer.Get("/commands", CommandsHandler)
	App.WebServer.Get("/whitelist", WhitelistHandler)
	//Post
	App.WebServer.Post("/commands", CmdHandler)
	App.WebServer.Post("/players", PlayerCmdHandler)
	App.WebServer.Post("/whitelist", WhitelistCmdHandler)

}
