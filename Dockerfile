FROM golang:1.12-stretch

LABEL maintainer="Rafael Nunes <rafaelnunes@engineer.com>"

RUN apt update && apt install sqlite3 -y
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/gorilla/mux

WORKDIR $GOPATH/src/github.com/rafaelcn/dataset-pipeline-server/

COPY . .

EXPOSE 4000
ENTRYPOINT ["make"]
