package handlers

import (
	"log/slog"

	"github.com/K1ender/questparty/internal/response"
	"github.com/K1ender/questparty/internal/storage"
	"github.com/K1ender/questparty/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type RoomHandler struct {
	roomStorage storage.RoomStorage
}

func NewRoomHandler(roomStorage storage.RoomStorage) *RoomHandler {
	return &RoomHandler{
		roomStorage: roomStorage,
	}
}

func (h *RoomHandler) JoinRoom(c fiber.Ctx) error {
	ctx := c.RequestCtx()

	err := h.roomStorage.JoinRoom(ctx, "", "")
	if err != nil {
		return response.NewInternalServerErrorResponse(c, "Internal Server Error")
	}

	return nil
}

type CreateRoomPayload struct {
	Username string `json:"username"`
}

func (h *RoomHandler) CreateRoom(c fiber.Ctx) error {
	ctx := c.RequestCtx()

	var payload CreateRoomPayload
	if err := c.Bind().Body(&payload); err != nil {
		slog.ErrorContext(ctx, "Error binding request body", "error", err)
		return response.NewBadRequestResponse(c, "Invalid request body")
	}

	slog.DebugContext(ctx, "Creating room", "username", payload.Username)

	id, err := utils.NanoID(8)
	slog.DebugContext(ctx, "Room ID", "id", id)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating room ID", "error", err)
		return response.NewInternalServerErrorResponse(c, "Internal Server Error")
	}

	err = h.roomStorage.CreateRoom(ctx, "", id)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating room", "error", err)
		return response.NewInternalServerErrorResponse(c, "Internal Server Error")
	}
	slog.DebugContext(ctx, "Room created", "id", id)

	err = h.roomStorage.JoinRoom(ctx, payload.Username, id)
	if err != nil {
		slog.ErrorContext(ctx, "Error joining room", "error", err)
		return response.NewInternalServerErrorResponse(c, "Internal Server Error")
	}

	slog.DebugContext(ctx, "Room joined", "username", payload.Username, "roomId", id)

	slog.DebugContext(ctx, "Room created", "id", id)
	return response.NewOkResponse(c, "Room created", map[string]string{"roomId": id})
}
