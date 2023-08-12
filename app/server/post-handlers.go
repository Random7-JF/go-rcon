package server

import (
	"fmt"
	"html/template"

	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/Random7-JF/go-rcon/app/validator"
	"github.com/gofiber/fiber/v2"
)

func PostCommandsHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessCmdForm(c)
	valid, err := SubmittedForm.CheckForReqFields()

	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		data["Response"] = err
		return c.Render("partials/response", td)
	}

	if valid {
		switch SubmittedForm.Cmd {
		case "say":
			rcon.SendMessage(AppConfig, SubmittedForm.Value)
		case "time":
			cmdresp, _ := rcon.SetTime(AppConfig, SubmittedForm.Value)
			data["Response"] = cmdresp
		case "weather":
			cmdresp := rcon.SetWeather(AppConfig, SubmittedForm.Options)
			data["Response"] = cmdresp
		case "kick":
			cmdresp, _ := rcon.KickPlayer(AppConfig, SubmittedForm.Value)
			data["Response"] = cmdresp
		case "setworldspawn":
			cmdresp := rcon.SetWorldSpawn(AppConfig, SubmittedForm.Value)
			data["Response"] = cmdresp
		default:
			data["Response"] = "No command found"
		}
	}

	return c.Render("partials/response", td)
}

func PostPlayersHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessCmdForm(c)
	valid, err := SubmittedForm.CheckForReqFields()
	data := make(map[string]interface{})

	td := model.TempalteData{
		Data: data,
	}

	if err != nil {
		data["Response"] = err
		return c.Render("partials/response", td)
	}

	players, err := rcon.GetPlayers(AppConfig)
	if err != nil {
		return err
	}
	data["Players"] = players

	if valid {
		switch SubmittedForm.Cmd {
		case "tp":
			//TODO implement
			cmdresp, _ := AppConfig.Rcon.Session.SendCommand(fmt.Sprintf("tp %s", SubmittedForm.Value))
			data["Response"] = cmdresp
		case "tpspawn":
			cmdresp, _ := rcon.TpToSpawn(AppConfig, SubmittedForm.Value)
			data["Response"] = cmdresp
		case "op":
			cmdresp, _ := AppConfig.Rcon.Session.SendCommand(fmt.Sprintf("op %s", SubmittedForm.Value))
			data["Response"] = cmdresp
		case "deop":
			cmdresp, _ := AppConfig.Rcon.Session.SendCommand(fmt.Sprintf("deop %s", SubmittedForm.Value))
			data["Response"] = cmdresp
		case "kick":
			cmdresp, _ := rcon.KickPlayer(AppConfig, SubmittedForm.Value)
			data["Response"] = cmdresp
		default:
			data["Response"] = "No command found"
		}
	}

	return c.Render("partials/response", td)
}

func PostWhiteListHandler(c *fiber.Ctx) error {

	SubmittedForm := validator.ProcessCmdForm(c)
	valid, err := SubmittedForm.CheckForReqFields()

	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if err != nil {
		data["Response"] = err
		return c.Render("partials/response", td)
	}

	if valid {
		cmdresp, _ := AppConfig.Rcon.Session.SendCommand(fmt.Sprintf("whitelist %s %s", SubmittedForm.Options, SubmittedForm.Value))
		data["Response"] = cmdresp
	}

	whitelist, err := rcon.GetWhitelist(AppConfig)
	if err != nil {
		fmt.Println("Error in whitelist: " + err.Error())
		data["Response"] = err
		template := template.Must(template.ParseFiles("views/pages/whitelist.html"))
		return template.ExecuteTemplate(c, "whitelist-table", td)
	}

	data["Whitelist"] = whitelist

	template := template.Must(template.ParseFiles("views/pages/whitelist.html"))
	return template.ExecuteTemplate(c, "whitelist-table", td)

}

func PostLoginHandler(c *fiber.Ctx) error {
	userForm := validator.ProcessUserForm(c)
	err := userForm.CheckForBlanks()

	if err != nil {
		helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
			Status:  false,
			Message: "Enter Username / Password",
		})
		return c.Redirect("/login")
	}

	err = model.Authenticate(userForm.User, userForm.Password)
	if err != nil {
		helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
			Status:  false,
			Message: "Incorrect Username / Password",
		})
		return c.Redirect("/login")
	}

	admin := model.IsUserAdmin(userForm.User)
	helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
		Status:  true,
		Message: "Successful Login",
		Admin:   admin,
	})

	return c.Redirect("/app/dashboard")
}

func PostUserHandler(c *fiber.Ctx) error {
	userForm := validator.ProcessUserForm(c)
	err := userForm.CheckForBlanks()
	if err != nil {
		fmt.Println("Form error: ", err)
		return c.Redirect("/app/admin/manage")
	}

	if userForm.Action == "create-user" {
		err = model.CreateUser(userForm.User, userForm.Password)
		if err != nil {
			fmt.Println("create user: ", err)
			return c.Redirect("/app/admin/manage")
		}
	} else if userForm.Action == "update-user-pass" {
		user, err := model.GetUserByUsername(userForm.User)
		if err != nil {
			fmt.Println("Get user error user: ", err)
			return c.Redirect("/app/admin/manage")
		}
		err = model.UpdateUserPass(int(user.ID), userForm.Password)
		if err != nil {
			fmt.Println("create user: ", err)
			return c.Redirect("/app/admin/manage")
		}
	} else if userForm.Action == "delete-user" {
		user, err := model.GetUserByUsername(userForm.User)
		if err != nil {
			fmt.Println("Get user error user: ", err)
			return c.Redirect("/app/admin/manage")
		}
		err = model.DeleteUser(user.ID)
		if err != nil {
			fmt.Println("Get user error user: ", err)
			return c.Redirect("/app/admin/manage")
		}
	}
	return c.Redirect("/app/admin/manage")
}
