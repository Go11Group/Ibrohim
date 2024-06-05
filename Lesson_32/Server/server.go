package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]struct{})
	clientsMu sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		clientsMu.Lock()
		clients[conn] = struct{}{}
		clientsMu.Unlock()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr().String())
	}()
	fmt.Println("Client connected:", conn.RemoteAddr().String())
	
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Println("Received message from", conn.LocalAddr().String() + ":", message)
		clientsMu.Lock()
		for c := range clients {
			if c != conn {
				_, err := c.Write([]byte("\n"+conn.LocalAddr().String() + ": "+message))
				if err != nil {
					fmt.Println("Error writing message to client:", err)
				}
			}
		}
		clientsMu.Unlock()
	}
}