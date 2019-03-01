CC=go
FLAGS=
SRC=database.go file.go main.go rest.go server.go uio.go
EXEC=server.out

all:
	go run $(SRC)

build:
	go build -o $(EXEC)

test:
	rm database/test.db -f
	go test -v *.go

debug:
	go run $(SRC) -v

clean:
	rm database/*.db
