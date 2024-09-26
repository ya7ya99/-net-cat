package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var (
	clients      = make(map[net.Conn]string)
	clientsMutex sync.Mutex
)

const maxClients = 10

var chathistory string

func main() {
	args := os.Args[1:]
	var port string

	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	} else if len(args) == 1 {
		port = args[0]
	} else {
		port = "8989"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port: " + port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleClient(conn)
	}
}
