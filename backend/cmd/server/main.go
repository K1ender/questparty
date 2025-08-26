package main

import (
	"net/http"

	"github.com/K1ender/questparty/internal/handlers"
	"github.com/K1ender/questparty/internal/storage"
)

func main() {
	inMemoryRoomStorage := storage.NewInMemoryStorage()
	roomHandlers := handlers.NewRoomHandler(inMemoryRoomStorage)

	http.HandleFunc("POST /api/join", roomHandlers.JoinRoom)
	http.HandleFunc("POST /api/create", roomHandlers.CreateRoom)

	http.ListenAndServe(":8080", nil)
}
