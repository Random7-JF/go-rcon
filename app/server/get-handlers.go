package server

import (
	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Home",
		Data:  data,
	}
	return c.Render("pages/index", td, "layouts/main")
}

func DashboardHandler(c *fiber.Ctx) error {
	players, err := rcon.GetPlayers()
	if err != nil {
		return err
	}
	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		return err
	}
	commands, err := model.GetCommandLog(5)
	if err != nil {
		return err
	}
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Players"] = players
	data["Rcon"] = AppConfig.RconSettings.Connection
	data["Whitelist"] = whitelist
	data["Commands"] = commands
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Dashboard",
		Data:  data,
	}

	return c.Render("pages/dashboard", td, "layouts/main")
}

func PlayersPageHandler(c *fiber.Ctx) error {
	players, err := rcon.GetPlayers()
	if err != nil {
		return err
	}
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Players"] = players
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Players",
		Data:  data,
	}

	return c.Render("pages/players", td, "layouts/main")
}

func CommandsHandler(c *fiber.Ctx) error {
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Commands",
		Data:  data,
	}
	return c.Render("pages/commands", td, "layouts/main")
}

func WhitelistHandler(c *fiber.Ctx) error {
	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		return err
	}
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Whitelist"] = whitelist
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Whitelist",
		Data:  data,
	}

	return c.Render("pages/whitelist", td, "layouts/main")
}

func LoginHandler(c *fiber.Ctx) error {

	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Login",
		Data:  data,
	}

	return c.Render("pages/login", td, "layouts/main")
}

func LogoutHandler(c *fiber.Ctx) error {
	helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
		Status:  false,
		Message: "Successful logout",
	})

	return c.Redirect("/")
}

func ManageHandler(c *fiber.Ctx) error {
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}
	commands, err := model.GetCommandLog(5)
	if err != nil {
		return err
	}
	users, err := model.GetAllUsers()
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Auth"] = auth
	data["Commands"] = commands
	data["Users"] = users

	td := model.TempalteData{
		Title: "Admin - Manage",
		Data:  data,
	}

	return c.Render("pages/admin/manage", td, "layouts/main")
}

func BenchHandler(c *fiber.Ctx) error {
	players, err := rcon.GetPlayers()
	if err != nil {
		return err
	}
	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		return err
	}
	commands, err := model.GetCommandLog(5)
	if err != nil {
		return err
	}
	auth, err := helper.GetAuthStatus(AppConfig, c)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Players"] = players
	data["Rcon"] = AppConfig.RconSettings.Connection
	data["Whitelist"] = whitelist
	data["Commands"] = commands
	data["Auth"] = auth

	td := model.TempalteData{
		Title: "Admin - Manage",
		Data:  data,
	}
	return c.Render("pages/bench", td, "layouts/main")

}
