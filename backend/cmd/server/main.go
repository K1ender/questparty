package main

import (
	"log/slog"

	"github.com/K1ender/questparty/internal/handlers"
	"github.com/K1ender/questparty/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	inMemoryRoomStorage := storage.NewInMemoryStorage()
	roomHandlers := handlers.NewRoomHandler(inMemoryRoomStorage)

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	app.Post("/api/join", roomHandlers.JoinRoom)
	app.Post("/api/create", roomHandlers.CreateRoom)

	app.Listen(":8080")
}
