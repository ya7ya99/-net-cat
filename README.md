# TCP Chat Application

This is a simple TCP-based chat application written in Go. It allows multiple clients to connect and chat with each other in real-time.

## Features

- Supports up to 10 simultaneous clients
- Unique usernames for each client
- Ability to change username during the chat
- Chat history stored in a file

## Requirements

- Go 1.15 or higher

## Installation

1. Clone the repository:
   ```
   https://learn.zone01oujda.ma/git/yakhaldy/net-cat
   ```
   
## Usage

### Starting the Server

To start the server, run the compiled binary with an optional port number:

```
./TCPChat [port]
```

If no port is specified, it will default to 8989.

### Connecting to the Server

Clients can connect to the server using any TCP client, such as netcat:

```
nc <localhost> <port>
```

Replace `<port>` with the port number the server is running on.

## Commands

- `-change_name <new_name>`: Change your username during the chat

## File Structure

- `main.go`: Contains the main server logic and client handling
- `NameExists.go`: Handles username validation and checking
- `ValidMessage.go`: Contains message validation logic
- `broadcastMessage.go`: Handles broadcasting messages to all connected clients
- `data.txt`: Stores the chat history

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.