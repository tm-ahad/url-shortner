package handlers

import (
	"context"
	"fmt"
	"os"
	"url-shortner/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ShortURL(client *redis.Client) fiber.Handler {
	cont := context.Background()

	return func(ctx *fiber.Ctx) error {
		url := string(ctx.Body())
		u, e := client.Get(cont, url).Result()

		host := os.Getenv("HOST")

		if e != nil {
			rng := helpers.RandString()

			shorten_url := fmt.Sprintf("http://%s/%s", host, rng)
			helpers.HandleErr(client.Set(cont, rng, url, 0).Err())
			helpers.HandleErr(client.Set(cont, url, rng, 0).Err())

			return ctx.SendString(shorten_url)
		} else {
			shorten_url := fmt.Sprintf("http://%s/%s", host, u)
			return ctx.SendString(shorten_url)
		}
	}
}
