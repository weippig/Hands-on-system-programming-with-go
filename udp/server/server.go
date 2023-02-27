package main

import (
	"bytes"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address. ")
	}

	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}

	conn, err := net.ListenUDP("udp", addr) // the UDP listener is a connection
	if err != nil {
		log.Fatalln("Listener:", os.Args[1], err)
	}

	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b) // ReceiveFrom method return the recipient's address
		if err != nil {
			log.Println("<-", addr, "Message error:", err)
			continue
		}

		msg := bytes.TrimSpace(b[:n])
		log.Printf("<- %q from %s", msg, addr)
		for i, l := 0, len(msg); i < l/2; i++ {
			msg[i], msg[l-1-i] = msg[l-1-i], msg[i]
		}

		msg = append(msg, '\n') // 加上換行符號
		if _, err := conn.WriteTo(b[:n], addr); err != nil {
			log.Println("->", addr, "Send error", err)
		}
	}
}
