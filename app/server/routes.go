package server

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func SetupRoutes(App *config.App) {
	//Public
	App.WebServer.Use(helmet.New())
	//Get
	App.WebServer.Get("/", IndexHandler)
	App.WebServer.Get("/login", LoginHandler)
	//Post
	App.WebServer.Post("/login", PostLoginHandler)

	//Protected
	user := App.WebServer.Group("/user")
	user.Use(
		func(c *fiber.Ctx) error {
			auth, err := helper.GetAuthStatus(AppConfig, c)
			if err != nil {
				fmt.Println("Middleware Auth Error:", err)
				return c.Redirect("/login")
			}
			if auth == nil {
				fmt.Println("Middleware Auth Error:", err)
				return c.Redirect("/login")
			}

			if auth.(model.Auth).Status {
				fmt.Println("Middleware Auth Success")
				return c.Next()
			}

			fmt.Println("Middleware Auth Error: Not Logged in")
			return c.Redirect("/")
		})
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

}
