package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn, port string) {
	defer conn.Close()

	// İstemciden gelen veriyi al
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Veri okuma hatası:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Printf("Sunucu %s:%s'den mesaj alındı: %s\n", conn.LocalAddr(), port, message)

	// Karşı sunucuya mesajı gönder
	sendMessageToOtherServer(port, message)
}
