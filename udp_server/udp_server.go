package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP Sunucu başlatıldı...")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Veri okuma hatası:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("İstemciden mesaj alındı: %s\n", message)

		// Mesajı büyük harfe çevirerek geri gönderelim
		response := fmt.Sprintf("Sunucu: %s", message)
		conn.WriteToUDP([]byte(response), clientAddr)
	}
}
