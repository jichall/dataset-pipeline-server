CC=go
FLAGS=
SRC=database.go file.go main.go rest.go server.go uio.go
EXEC=server

all:
	go run $(SRC)

build:
	go build -o $(EXEC)

test:
	go test -v *.go

clean:
	rm database/*.db
