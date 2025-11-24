# 💬 RPC Chat Room (GoLang Project)

A simple **chatroom application** built using Go's **net/rpc** package.
Clients can connect to a central RPC server, send messages, and fetch full chat history.

---

## 🚀 Features

* 📡 Client–Server communication using **Go RPC**
* 💾 Server stores all messages in memory
* 👥 Multiple clients can send messages (chat history shared)
* 🔁 Each message returns the full chat history
* 🧹 Simple and clean console interface

---

## 🧱 Project Structure

```
RPC CHAT ROOM/
│
├── client.go      # Client-side program
├── server.go      # RPC server
├── go.mod         # Go module
└── README.md      # Project documentation
```

---

## ⚙️ How to Run

### 1️⃣ Run the Server

Open a terminal inside the project folder:

```bash
go run server.go
```

Server will start on port `:1234` and print incoming messages.

---

### 2️⃣ Run the Client

open another terminal (many as you want):

```bash
go run client.go
```

* Enter your **name** when prompted
* Type messages — they’ll be sent to the server
* The **entire chat history** will display after each message
* Type `exit` to leave

---

## 🖼 Example Output

**Server terminal:**

```
Chat server running on port 1234...
Ahmed: Hello!
Omar: Hi Ahmed, how are you?
```

**Client terminal:**

```
Enter your name: Ahmed
Welcome Ahmed! You've joined the chat.
Enter message (or 'exit' to quit): Hello!
--- Chat History ---
Ahmed: Hello!
--------------------
```

---

## 🧩 Technologies Used

* [GoLang](https://go.dev/)
* net/rpc package
* bufio, fmt, log, strings, sync

---

## 📜 License

This project is licensed under the **MIT License** – see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Ahmed Elsafty**
📧 [elsaftyahmed09@gmail.com]

⭐ Feel free to fork, improve, and star this repo!
