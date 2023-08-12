package middleware

import (
	"errors"
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/helper"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/gofiber/fiber/v2"
)

type Mwconfig struct {
	AppConfig *config.App
}

func (mw Mwconfig) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth, err := helper.GetKey(mw.AppConfig, c, "Auth")
		if err != nil {
			return c.Redirect("/login")
		}
		if auth == nil {
			fmt.Println("Middleware Auth Error:", err)
			return c.Redirect("/login")
		}

		if auth.(model.Auth).Status {
			return c.Next()
		}
		return c.Redirect("/login")
	}
}

func (mw Mwconfig) SaveSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := mw.AppConfig.Store.Get(c)
		if err != nil {
			return errors.New("unable to get session store")
		}
		if err := session.Save(); err != nil {
			return err
		}
		return c.Next()
	}
}

func (mw Mwconfig) SetupSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := mw.AppConfig.Store.Get(c)
		if err != nil {
			return errors.New("unable to get session store")
		}

		auth := session.Get("Auth")
		if auth == nil {
			helper.UpdateSessionKey(mw.AppConfig, c, "Auth", model.Auth{Status: false, Message: "", Admin: false})
		}

		return c.Next()
	}
}
