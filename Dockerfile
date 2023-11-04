FROM golang:1.21.3

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
