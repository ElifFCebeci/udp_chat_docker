# Go 1.20-alpine imajını temel al
FROM golang:1.20-alpine

# Çalışma dizinini /app olarak ayarla
WORKDIR /app

# udp_server.go dosyasını Docker konteynerine kopyala
COPY udp_server.go .

# udp_server.go dosyasını derle ve udp_server adında bir çalıştırılabilir dosya oluştur
RUN go build -o udp_server udp_server.go

# Konteyner başladığında udp_server çalıştırılabilir dosyasını çalıştır
CMD ["./udp_server"]
