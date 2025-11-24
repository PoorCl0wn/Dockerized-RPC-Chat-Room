package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name string
	Text string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		var reply []string
		err = client.Call("ChatServer.SendMessage", Message{Name: name, Text: text}, &reply)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range reply {
			fmt.Println(msg)
		}
		fmt.Println("--------------------")
	}
}
