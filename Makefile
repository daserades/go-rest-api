run:
	sudo go run ./cmd/go-rest-api/main.go
	clear

build:
	sudo go build -ldflags "-s -w" -o service
	clear
	./service

buildw:
	go build -ldflags "-s -w" -o service.exe
	./service.exe