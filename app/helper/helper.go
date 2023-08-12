package helper

import (
	"errors"
	"log"

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

func GetKey(app *config.App, c *fiber.Ctx, key string) (interface{}, error) {
	session, err := app.Store.Get(c)
	if err != nil {
		return nil, errors.New("unable to get session store")
	}
	reqKey := session.Get(key)
	log.Printf("%s Key Data: %v", key, reqKey)
	return reqKey, nil
}
