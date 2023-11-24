package handlers

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func Redirect(client *redis.Client) fiber.Handler {
	cont := context.Background()

	return func(ctx *fiber.Ctx) error {
		eurl := ctx.Params("id")

		u, e := client.Get(cont, eurl).Result()
		
		if len(eurl) == 0 {
			u = "/static/short_url.html"
		} else if e != nil {
			u = os.Getenv("NOT_FOUND_PAGE")
		}

		return ctx.Redirect(u)
	}
}
