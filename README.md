# Dockerized RPC Chatroom (Go)

A simple Go RPC chatroom packaged in a Docker container. Run the server locally, test it with the Go client, and pull the Docker image from Docker Hub.

## Docker Hub Image
[https://hub.docker.com/r/poorclown/rpc-chat-server](https://hub.docker.com/r/poorclown/rpc-chat-server)

## How to Run

### Run the Server (Docker)
```bash
docker pull poorclown/rpc-chat-server:v1
docker run --rm -p 1234:1234 poorclown/rpc-chat-server:v1
```

### Run the Client (Go)
```bash
go run client.go
```
The client connects to the server at `localhost:1234`.
