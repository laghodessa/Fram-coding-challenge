# build stage
FROM docker.io/golang:1.18.3 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -ldflags "-linkmode external -extldflags -static" ./cmd/server

FROM docker.io/alpine:3.16.0

WORKDIR /app
COPY --from=builder /app/server server

USER nobody
ENTRYPOINT [ "./server" ]
