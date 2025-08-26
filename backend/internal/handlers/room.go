package handlers

import (
	"net/http"

	"github.com/K1ender/questparty/internal/storage"
)

type RoomHandler struct {
	roomStorage storage.RoomStorage
}

func NewRoomHandler(roomStorage storage.RoomStorage) *RoomHandler {
	return &RoomHandler{
		roomStorage: roomStorage,
	}
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.roomStorage.JoinRoom(ctx, "", "")
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.roomStorage.CreateRoom(ctx, "")
}
