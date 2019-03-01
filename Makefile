CC=go
FLAGS=
SRC=database.go file.go main.go rest.go server.go uio.go
EXEC=server.out

all:
	cd src && go run $(FLAGS) $(SRC)

build:
	cd src && go build -o $(EXEC)

test:
	rm database/test.db -f
	cd src && go test -v *.go

debug:
	cd src && go run $(SRC) -v

ci:
	docker run --entrypoint=cd src && go test -v

clean:
	rm database/*.db
