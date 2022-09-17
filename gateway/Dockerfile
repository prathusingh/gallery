FROM golang:1.19.1

ENV config=docker

WORKDIR /app

COPY .. /app

RUN go get github.com/github/CompileDaemon
RUN go mod download

ENTRYPOINT CompileDaemon --build="go build -o main gateway/cmd/main.go" --command=./main