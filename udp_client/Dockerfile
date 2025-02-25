# Küçük boyutlu Go imajını kullan
FROM golang:1.20-alpine

# Çalışma dizinini belirle
WORKDIR /app

# Kod dosyasını kopyala
COPY udp_client.go .

# Go kodunu derle
RUN go build -o udp_client udp_client.go

# Çalıştırılacak komut
CMD ["./udp_client"]
