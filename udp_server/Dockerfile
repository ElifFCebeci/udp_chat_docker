FROM golang:1.20-alpine

WORKDIR /app

COPY udp_server.go .

RUN go build -o udp_server udp_server.go

CMD ["./udp_server"]
