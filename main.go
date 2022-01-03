package main

import (
	"filme-api/servicos"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/movies/:id?", servicos.BuscarMovie)
	app.Post("/movies", servicos.AddMovie)
	app.Delete("/movies/:id", servicos.DelMovie)

	app.Listen(":8000")
}
