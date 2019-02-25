CC=go
FLAGS=
SRC=database.go file.go main.go rest.go server.go uio.go

all:
	go run $(SRC)

test:
	go test -v *.go

clean:
	rm database/*.db
