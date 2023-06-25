package helper

import (
	"errors"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/gofiber/fiber/v2"
)

func UpdateSessionKey(app *config.App, c *fiber.Ctx, key string, value interface{}) error {
	session, err := app.Store.Get(c)
	if err != nil {
		return errors.New("unable to get session store")
	}
	session.Set(key, value)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func GetAuthStatus(app *config.App, c *fiber.Ctx) (interface{}, error) {
	session, err := app.Store.Get(c)
	if err != nil {
		return nil, errors.New("unable to get session store")
	}
	auth := session.Get("Auth")
	return auth, nil
}
