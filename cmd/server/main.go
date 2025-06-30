package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"

	"go-vpn/internal/tlsconfig"
)

func main() {
	config, err := tlsconfig.LoadTLSConfig("cmd/cert/server.crt", "cmd/cert/server.key", "cmd/cert/ca.crt", true) // сервер

	if err != nil {
		log.Fatal(err)
	}

	listener, err := tls.Listen("tcp", ":4433", config)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("VPN Server is running on port 4433...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept failed:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("New TLS connection from", conn.RemoteAddr())
	io.Copy(conn, conn) // эхо-сервер
}
