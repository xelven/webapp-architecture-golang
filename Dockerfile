# Build environment
# -----------------
FROM golang:1.21-alpine as build-env
WORKDIR /webapp-core-1

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api

# Deployment environment
# ----------------------
FROM alpine

COPY --from=build-env /webapp-core-1/bin/api /webapp-core-1/

EXPOSE 8080
CMD ["/webapp-core-1/api"]
