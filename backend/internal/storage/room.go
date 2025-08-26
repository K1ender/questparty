package storage

import (
	"context"
)

type RoomStorage interface {
	JoinRoom(ctx context.Context, username string, roomId string) error
	CreateRoom(ctx context.Context, username string, roomId string) error
}

type GameState struct {
	Users []string
}

type InMemoryStorage struct {
	Rooms map[string]GameState
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{Rooms: make(map[string]GameState)}
}

func (s *InMemoryStorage) JoinRoom(ctx context.Context, username string, roomId string) error {
	room := s.Rooms[roomId]
	room.Users = append(room.Users, username)
	s.Rooms[roomId] = room
	return nil
}

func (s *InMemoryStorage) CreateRoom(ctx context.Context, username string, roomId string) error {
	s.Rooms[roomId] = GameState{}
	return nil
}
