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
	App.WebServer.Get("/user/login", LoginHandler)
	//Post
	App.WebServer.Post("/commands", PostCommandsHandler)
	App.WebServer.Post("/players", PostPlayersHandler)
	App.WebServer.Post("/whitelist", PostWhitelistHandler)
	App.WebServer.Post("/user/login", PostLoginHandler)

}
