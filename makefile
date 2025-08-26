dev:
	cd frontend && bun run dev &
	go run ./backend/cmd/server/main.go