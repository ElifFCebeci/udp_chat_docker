package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Received: %s\n", string(buffer[:n]))
		}
	}()

	for {
		var msg string
		fmt.Print("Mesaj girin: ")
		fmt.Scanln(&msg)

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}
