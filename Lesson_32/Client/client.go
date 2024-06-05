package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server.")

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			serverMessage, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Println(serverMessage)
		}
	}()

	userReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		userInput, err := userReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		fmt.Fprint(conn, strings.TrimSpace(userInput) + "\n")
	}
	fmt.Println("Connection closed.")
}