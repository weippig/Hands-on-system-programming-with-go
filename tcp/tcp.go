package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 執行範例： go run main.go localhost:8080
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address.")
	}

	// 將 IP + port 解析為 TCPaddr
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}

	// 建立 TCP Listener
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln("Listener:", os.Args[1], err)
	}
	log.Println("<- Listening on", addr)

	go createConn(addr)

	conn, err := listener.AcceptTCP()
	if err != nil {
		log.Fatalln("<- Accrpt:", os.Args[1], err)
	}
	handleConn(conn)
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

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	time.Sleep(time.Second / 2)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println("<- Message error:", err)
			continue
		}

		switch msg = strings.TrimSpace(msg); msg {
		case `\q`: // 輸入 \q 的話會離開
			log.Println("Exiting...")
			if err := conn.Close(); err != nil {
				log.Println("<- Close:", err)
			}
			time.Sleep(time.Second / 2)
			return
		case `\x`: // 輸入 \x 的話 會顯示下面的字
			log.Println("<- Special message `\\x` received!")
		default:
			log.Println("<- Message Received:", msg)
		}
	}
}
