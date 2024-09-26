package main

import (
	"fmt"
	"net"
)

func broadcastMessage(message string, cl net.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for conn := range clients {
		if cl != conn {
			_, err := conn.Write([]byte(message))

			Writemsg(conn)

			if err != nil {
				fmt.Printf("Error broadcasting to client: %v\n", err)
				conn.Close()
				delete(clients, conn)
			}
		}
	}
}
