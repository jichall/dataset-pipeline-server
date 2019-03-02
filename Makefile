CC=go
FLAGS=
SRC=database.go file.go main.go rest.go server.go uio.go
EXEC=server.out

all:
	cd src && go run $(FLAGS) $(SRC)

build:
	cd src && go build -o $(EXEC)

deploy:
	cd src && go run $(FLAGS) $(SRC) -d

test:
	rm database/test.db -f
	cd src && go test -v *.go

debug:
	cd src && go run $(SRC) -v

docker_ci:
	docker build -t dps .
	docker run -t --entrypoint=make dps test

docker_run:
	docker run -t dps

clean:
	rm database/*.db
