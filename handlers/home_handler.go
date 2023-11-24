package handlers

import "github.com/gofiber/fiber/v2"

func Home(ctx *fiber.Ctx) error {
	return ctx.Redirect("/static/short_url.html")
}
