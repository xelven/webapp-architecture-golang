FROM golang:1.21-alpine
WORKDIR /webapp-core

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 8080

# ENTRYPOINT CompileDaemon --build="go build ./cmd/api/main.go" --command="./main"
