CC=go
FLAGS=
SRC=src/database.go src/file.go src/main.go src/rest.go src/server.go src/uio.go

all:
	go run $(SRC)

test:
	go test src/*.go
