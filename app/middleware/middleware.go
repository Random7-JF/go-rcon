package middleware

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/gofiber/fiber/v2"
)

type Mwconfig struct {
	AppConfig *config.App
}

func (m Mwconfig) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth, err := helper.GetAuthStatus(m.AppConfig, c)
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
		return c.Next()
	}
}
