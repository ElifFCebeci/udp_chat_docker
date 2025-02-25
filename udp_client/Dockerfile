# Küçük boyutlu Go imajı
FROM golang:1.20-alpine

# Çalışma dizini
WORKDIR /app

# Kod dosyasını kopyalama
COPY udp_client.go .

# Go kodunu derleme
RUN go build -o udp_client udp_client.go

# Çalıştırılacak komut
CMD ["./udp_client"]
