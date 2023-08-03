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
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/commands")
	}

	data := make(map[string]interface{})
	data["Response"] = "Sent Message"

	td := model.TempalteData{
		Data: data,
	}

	return c.Render("partials/response", td)
}

func PostPlayersHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()
	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/players")
	}

	return c.Redirect("/app/players")
}

func PostWhitelistHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/whitelist")
	}

	return c.Redirect("/app/whitelist")

}

func PostWhiteListHandler(c *fiber.Ctx) error {
	SubmittedForm := validator.ProcessForm(c)
	err := SubmittedForm.ValidateInputs()

	if err != nil {
		fmt.Println("Error in form submission: " + err.Error())
		return c.Redirect("/app/whitelist")
	}

	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		fmt.Println("Error in whitelist: " + err.Error())
		return c.Redirect("/app/whitelist")
	}

	data := make(map[string]interface{})
	data["Whitelist"] = whitelist

	td := model.TempalteData{
		Data: data,
	}

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

	helper.UpdateSessionKey(AppConfig, c, "Auth", model.Auth{
		Status:  true,
		Message: "Successful Login",
		Admin:   model.IsUserAdmin(userForm.User),
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
