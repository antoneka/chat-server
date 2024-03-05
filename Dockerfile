FROM golang:1.22.0-alpine AS builder

COPY . /github.com/antoneka/chat-server/source
WORKDIR /github.com/antoneka/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/antoneka/chat-server/source/bin/chat_server .
COPY .env .

CMD ["./chat_server"]