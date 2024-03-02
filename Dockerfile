FROM golang:1.22.00-alpine AS builder

COPY . /github.com/antoneka/chat-server/source
WORKDIR /github.com/antoneka/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat-server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/antoneka/chat-server/source/bin/chat-server .

CMD ["./chat-server"]