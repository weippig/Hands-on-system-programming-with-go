package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address.")
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}
	createConn(addr)
}

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing")
	// 建立 connection
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln("-> Connection:", err)
	}
	log.Println("-> Connection to", addr)

	// 從標準輸入 read messages
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("# ")
		msg, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatalln("-> Message error:", err)
		}

		// 再將 messages 轉送給 connection
		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connection", err)
			return
		}
	}
}
