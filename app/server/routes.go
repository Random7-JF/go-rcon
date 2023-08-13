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
	App.WebServer.Use(mw.SetupSession())
	//TODO make this work - need to do more research
	//App.WebServer.Use(csrf.New())
	//Get
	App.WebServer.Get("/", IndexHandler)
	App.WebServer.Get("/bench", BenchHandler)
	App.WebServer.Get("/login", LoginHandler)
	//Post
	App.WebServer.Post("/login", PostLoginHandler)

	//Protected
	user := App.WebServer.Group("/app", mw.Auth(), mw.SaveSession())
	// Get
	user.Get("/dashboard", DashboardHandler)
	user.Get("/players", PlayersPageHandler)
	user.Get("/commands", CommandsHandler)
	user.Get("/whitelist", WhitelistHandler)
	user.Get("/logout", LogoutHandler)
	//HTMX Endpoints
	user.Get("/players/list", PlayerListHandler)
	user.Get("/players/count", PlayerCountHandler)
	user.Get("/commands/list", CommandsListHandler)

	// Post
	user.Post("/commands", PostCommandsHandler)
	user.Post("/players", PostPlayersHandler)
	user.Post("/login", PostLoginHandler)
	//HTMX Endpoints
	user.Post("/whitelist/update", PostWhiteListHandler)

	admin := user.Group("/admin", mw.Auth(), mw.SaveSession())
	admin.Get("/manage", ManageHandler)
	admin.Get("/metrics", monitor.New())
	admin.Post("/user/update", PostUserUpdate)
	admin.Post("/user/remove", PostUserRemove)
	admin.Post("/user", PostUserHandler)
	admin.Post("/rcon", PostRconHandler)
	admin.Post("/rcon/session", PostRconSessionHandler)

}
