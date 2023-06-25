package server

import (
	"time"

	"github.com/Random7-JF/go-rcon/app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

var AppConfig *config.App

func Serve(App *config.App) {
	// Create a new engine
	htmlEngine := html.New("./views", ".html")

	htmlEngine.Debug(!App.Production)
	htmlEngine.Reload(!App.Production)

	App.WebServer = fiber.New(fiber.Config{
		Views: htmlEngine,
	})

	App.Store = session.New(session.Config{Expiration: 10 * time.Minute})
	App.WebServer.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	App.WebServer.Static("/static", "./views/static")

	SetupRoutes(App)
	AppConfig = App

	url := App.WebSettings.Ip + ":" + App.WebSettings.Port
	App.WebServer.Listen(url)
}
