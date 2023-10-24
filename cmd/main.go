package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/immersivesky/go-rest-postgresql/internal/adapters/handlers"
	"github.com/immersivesky/go-rest-postgresql/internal/adapters/repository"
)

func main() {
	app := fiber.New()

	db := repository.NewDB("postgres://postgresuser:postgrespass@localhost:5678/vk")
	handlers := handlers.NewHTTPHandlers(db)

	app.Get("/chat/:id", handlers.GetChat)
	app.Post("/chat/:id", handlers.CreateChat)

	log.Fatalln(app.Listen(":3200"))
}
