package server

import (
	"fmt"
	"html/template"
	"log"

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

func PostUserUpdate(c *fiber.Ctx) error {
	action := c.FormValue("action")
	value := c.FormValue("value")
	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if len(action) == 0 || len(value) == 0 {
		data["Response"] = "Error: Blank submission"
		log.Println("PostUserUpdate : Blank submission ")
		return c.Render("partials/response", td)
	}

	if action == "admin-user" {
		err := model.SetUserAdmin(value, true)
		if err != nil {
			data["Response"] = err.Error()
			log.Println("PostUserUpdate - SetUserAdmin: ", err.Error())
			return c.Render("partials/response", td)

		}
		data["Response"] = "User set as admin."

	} else if action == "remove-admin-user" {
		err := model.SetUserAdmin(value, false)
		if err != nil {
			data["Response"] = err.Error()
			log.Println("PostUserUpdate - SetUserAdmin: ", err.Error())
			return c.Render("partials/response", td)
		}
		data["Response"] = "User no longer admin."
	}

	return c.Render("partials/response", td)
}

func PostUserRemove(c *fiber.Ctx) error {
	target := c.FormValue("value")
	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if len(target) == 0 {
		data["Response"] = "Error: Blank submission"
		log.Println("PostUserRemove : Blank submission ")
		return c.Render("partials/response", td)
	}

	user, err := model.GetUserByUsername(target)
	if err != nil {
		data["Response"] = "GetUserByUsername Error: " + err.Error()
		log.Println("PostUserRemove - GetUserByUsername: " + err.Error())
		return c.Render("partials/response", td)

	}
	err = model.DeleteUser(user.ID)
	if err != nil {
		data["Response"] = "DeleteUser Error: " + err.Error()
		log.Println("PostUserRemove - DeleteUser: " + err.Error())
		return c.Render("partials/response", td)

	}

	data["Response"] = "User deleted."
	return c.Render("partials/response", td)
}

func PostRconHandler(c *fiber.Ctx) error {

	rconForm := validator.ProcessRconForm(c)
	err := rconForm.CheckForReqFields()
	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if err != nil {
		log.Println("PostRconHandler - CheckForReqFields: " + err.Error())
		data["Response"] = "PostRconHandler - CheckForReqFields: " + err.Error()
		return c.Render("partials/response", td)

	}

	err = model.SetRconSettings(rconForm.Ip, rconForm.Port, rconForm.Password)
	if err != nil {
		log.Println("PostRconHandler - SetRconSettings: " + err.Error())
		data["Response"] = "PostRconHandler - SetRconSettings: " + err.Error()
		return c.Render("partials/response", td)

	}

	data["Response"] = "Rcon Settings Updated."

	return c.Render("partials/response", td)
}

func PostRconSessionHandler(c *fiber.Ctx) error {
	rconSessionForm := validator.ProcessRconSessionForm(c)
	err := rconSessionForm.CheckForReqFields()
	data := make(map[string]interface{})
	td := model.TempalteData{
		Data: data,
	}

	if err != nil {
		log.Println("PostRconSessionHandler - CheckForReqFields: " + err.Error())
	}

	switch rconSessionForm.Action {
	case "stop":
		rcon.DisconnectSession(AppConfig)
		data["Response"] = "Rcon Disconnected"
	case "start":
		err = rcon.ConnectSession(AppConfig)
		if err != nil {
			data["Response"] = err
			return c.Render("partials/response", td)
		}
		data["Response"] = "Rcon Connected"

	case "restart":
		if !AppConfig.Rcon.Connection {
			err = rcon.ConnectSession(AppConfig)
			if err != nil {
				data["Response"] = err
				return c.Render("partials/response", td)
			}
		}

		rcon.DisconnectSession(AppConfig)
		err = rcon.ConnectSession(AppConfig)
		if err != nil {
			data["Response"] = err
			return c.Render("partials/response", td)
		}
		data["Response"] = "Rcon Reconnected"

	}

	return c.Render("partials/response", td)
}
