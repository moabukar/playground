# TCP

## How a TCP Server Works

- Basics of TCP: TCP stands for Transmission Control Protocol. Think of it like making a phone call. You dial, the other person picks up, and you have a continuous, two-way conversation. TCP establishes a connection and maintains it for continuous, reliable communication.

- TCP Server Role: A TCP server is like a phone operator who is ready to take calls. It listens for incoming connections (calls), and when a client (caller) connects, it establishes a line of communication.

- Reliable Communication: Unlike UDP, TCP checks to make sure all the data (like words in your conversation) arrives correctly and in order. If something goes wrong, it asks for the data to be sent again.

## What's Happening in the Example


- Listening for Connections: The server starts by listening for incoming connections on a specific port (":8080"). It's like our operator sitting by the phone, waiting for a call.

- Accepting Connections: When a client tries to connect, the server accepts the connection. This is like our operator answering the phone.

- Handling the Connection: Each connection is handled separately. The handleConnection function is like a dedicated line where the operator can talk to one caller at a time. By using go handleConnection(conn), each connection is handled concurrently (like having multiple operators for multiple calls).

- Reading Messages: The server reads messages sent by the client. In TCP, messages are sent in a stream, so it reads until it finds a newline character ('\n'). This is like listening to what the caller is saying until they pause.

- Printing the Message: After receiving a message, the server prints it out. This is like the operator writing down the caller's message.

- Continuous Operation: The server continues to listen for new connections and handle them, just like an operator who's always ready to take the next call.

- In our test, you'd use a TCP client to connect to the server (dial the number) and then send a message (speak into the phone). The server (operator) listens, accepts the connection (picks up the phone), and then reads and displays the message (listens and notes down what you say).

- The key difference from the UDP server is the connection setup and maintenance, ensuring reliable communication. It's like having a conversation on a reliable phone line, compared to sending letters via mail (UDP) where you don't know if they arrive

## Test the server (via telnet or netcat)

- `go run tcpserver.go`

- `telnet localhost 8083`
- `nc localhost 8083`
  -  `echo "Hello, TCP Server" | nc localhost 8083`
