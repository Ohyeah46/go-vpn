package main

import (
	"crypto/tls"
	"io"
	"log"

	"go-vpn/internal/tlsconfig"
)

func main() {
	log.Println("Starting VPN client...")
	config, err := tlsconfig.LoadTLSConfig("cmd/cert/client.crt", "cmd/cert/client.key", "cmd/cert/ca.crt", false) // клиент
	if err != nil {
		log.Fatal("Loading TLS config failed:", err)
	}

	conn, err := tls.Dial("tcp", "localhost:4433", config)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	log.Println("Connected to server.")

	io.Copy(conn, conn)
}
