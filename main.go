package main

import (
	"os"
	"url-shortner/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	app := fiber.New()
	godotenv.Load("./config/.env")

	host := os.Getenv("HOST")

	var client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_HOST"),
		Password: "",
		DB: 0,
	})

	app.Post("/short_url", handlers.ShortURL(client))
	app.Get("/:id", handlers.Redirect(client))
	app.Static("/static", "./static")
	app.Get("/", handlers.Home)

	app.Listen(host);
}
