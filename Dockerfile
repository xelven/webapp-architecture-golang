FROM golang:1.21-alpine
WORKDIR /webapp-core

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api
EXPOSE 8080
