package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"shopinbiz.telegram.bot/clients"
	"shopinbiz.telegram.bot/config"
	"shopinbiz.telegram.bot/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	bot := clients.Init()
	handlers.Init(bot)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
