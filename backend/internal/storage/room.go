package storage

import "context"

type RoomStorage interface {
	JoinRoom(ctx context.Context, username string, roomId string) error
	CreateRoom(ctx context.Context, username string) (string, error)
}

type idk any

type InMemoryStorage struct {
	Rooms map[string]idk
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{Rooms: make(map[string]idk)}
}

func (s *InMemoryStorage) JoinRoom(ctx context.Context, username string, roomId string) error {
	s.Rooms[roomId] = username
	return nil
}

func (s *InMemoryStorage) CreateRoom(ctx context.Context, username string) (string, error) {
	return "", nil
}
