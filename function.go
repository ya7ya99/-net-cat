package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func WriteData(chat string) {
	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(chat))
	if err != nil {
		fmt.Println(err)
	}
}

func Writemsg(conn net.Conn) {
	currentTime := time.Now().Format("[2006-01-02 15:04:05]")
	formattedMessage := fmt.Sprintf(Green+"%s["+res+Blue+"%s"+res+Green+"]: "+res, currentTime, clients[conn])
	conn.Write([]byte(formattedMessage))
}
