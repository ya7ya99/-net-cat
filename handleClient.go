package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var (
	Green      string = "\033[92m"
	DarkYellow string = "\033[33m"
	Blue       string = "\033[94m"
	res        string = "\033[00m"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientsMutex.Lock()
	if len(clients) >= maxClients {
		clientsMutex.Unlock()
		conn.Write([]byte("Sorry, the chat is full. Try again later.\n"))
		conn.Close()
		return
	}
	clientsMutex.Unlock()

	conn.Write([]byte(Green + "Welcome to TCP-Chat! \n" + res))
	conn.Write([]byte(

		DarkYellow + "         _nnnn_\n" +
			"        dGGGGMMb\n" +
			"       @p~qp~~qMb\n" +
			"       M|@||@) M|\n" +
			"       @,----.JM|\n" +
			"      JS^\\__/  qKL\n" +
			"     dZP        qKRb\n" +
			"    dZP          qKKb\n" +
			"   fZP            SMMb\n" +
			"   HZM            MMMM\n" +
			"   FqM            MMMM\n" +
			" __| \".        |\\dS\"qML\n" +
			" |    `.       | `' \\Zq\n" +
			"_)      \\.___.,|     .'\n" +
			"\\____   )MMMMMP|   .'\n" +
			"     `-'       `--'\n" + res,
	))
	var nameStr string
	scanner := bufio.NewScanner(conn)
	for {
		conn.Write([]byte(Green + "[ENTER YOUR NAME]: " + res))
		scanner.Scan()
		nameStr = scanner.Text()
		if len(nameStr) > 0 && NameExists(nameStr) {
			break
		}
		conn.Write([]byte("Invalid or duplicate name. Try again.\n"))
	}

	clientsMutex.Lock()
	clients[conn] = nameStr
	clientsMutex.Unlock()

	WriteData("\n" + nameStr + " has joined our chat...\n")
	broadcastMessage("\n"+nameStr+" has joined our chat...\n", conn)

	conn.Write([]byte(chathistory))

	Writemsg(conn)

	for scanner.Scan() {

		Writemsg(conn)

		messageStr := scanner.Text()
		if len(messageStr) > 1000 {
			conn.Write([]byte("\nthe lenght of the message cannot being more than 1000 caracter\n"))
			Writemsg(conn)
		} else {
			currentTime := time.Now().Format("[2006-01-02 15:04:05]")
			formattedMessage := fmt.Sprintf("\033[92m%s[\033[0m\033[94m%s\033[0m\033[92m]\033[0m: %s\n", currentTime, clients[conn], messageStr)
			flag, err := ChackFlag(messageStr)
			if flag != "" {
				if err != "" {
					conn.Write([]byte(err))
				} else {
					clientsMutex.Lock()
					clients[conn] = flag
					clientsMutex.Unlock()
					WriteData("\n" + nameStr + " has change their name to " + flag + "\n")
					broadcastMessage("\n"+nameStr+" has change their name to "+flag+"\n", conn)
					conn.Write([]byte("\n" + nameStr + " has change their name to " + flag + "\n"))
					Writemsg(conn)
				}
			} else if ValidMessage(messageStr) {
				chathistory += formattedMessage
				WriteData(fmt.Sprintf("%s[%s]: %s\n", currentTime, clients[conn], messageStr))
				broadcastMessage("\n"+formattedMessage, conn)
			}
		}

	}
	WriteData(clients[conn] + " has left our chat...\n")
	broadcastMessage("\n"+clients[conn]+" has left our chat...\n", conn)
	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()
}
