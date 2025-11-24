package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// Message holds a user's name and message text
type Message struct {
	Name string
	Text string
}

// ChatServer keeps chat history
type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

// SendMessage adds a message to the chat and returns history
func (cs *ChatServer) SendMessage(msg Message, reply *[]string) error {
	cs.mu.Lock()
	formatted := fmt.Sprintf("%s: %s", msg.Name, msg.Text)
	cs.messages = append(cs.messages, formatted)
	*reply = cs.messages
	fmt.Println(formatted) // print on server terminal
	cs.mu.Unlock()
	return nil
}

// GetMessages returns the full chat history
func (cs *ChatServer) GetMessages(dummy bool, reply *[]string) error {
	cs.mu.Lock()
	*reply = append([]string(nil), cs.messages...)
	cs.mu.Unlock()
	return nil
}

func main() {
	chatServer := new(ChatServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	fmt.Println("Chat server running on port 1234...")
	rpc.Accept(listener)
}
