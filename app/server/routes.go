package server

import (
	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/middleware"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(App *config.App) {
	mw := middleware.Mwconfig{AppConfig: App}
	//Public
	App.WebServer.Use(helmet.New())
	//Get
	App.WebServer.Get("/", IndexHandler)
	App.WebServer.Get("/login", LoginHandler)
	//Post
	App.WebServer.Post("/login", PostLoginHandler)

	//Protected
	user := App.WebServer.Group("/user", mw.Auth())
	// Get
	user.Get("/dashboard", DashboardHandler)
	user.Get("/players", PlayersPageHandler)
	user.Get("/commands", CommandsHandler)
	user.Get("/whitelist", WhitelistHandler)
	user.Get("/logout", LogoutHandler)
	// Post
	user.Post("/commands", PostCommandsHandler)
	user.Post("/players", PostPlayersHandler)
	user.Post("/whitelist", PostWhitelistHandler)
	user.Post("/login", PostLoginHandler)

	admin := App.WebServer.Group("/admin", mw.Auth())

	admin.Get("/metrics", monitor.New())

}
