# UDP

## How a UDP Server Works

- Basics of UDP: UDP stands for User Datagram Protocol. It's like sending a letter through the mail. You send it off (as a packet of data), but you don't get a confirmation that it was delivered. It's quick because there's no back-and-forth communication required to start the conversation.

- UDP Server Role: A UDP server is like a mailbox that's ready to receive these letters (data packets). It doesn't establish a continuous connection with the sender. Instead, it just listens for any packets sent its way and processes them.

- No Connection Establishment: Unlike TCP, UDP doesn't establish a connection before data transfer. This means less time setting up and more time doing.

## Example UDP in Go


- Setting Up an Address: The server starts by setting up an address (":8080"). Think of this as deciding the mailbox number where it's going to receive letters.

- Listening for Packets: The server then starts listening for any UDP packets sent to this address. It's like having someone constantly checking the mailbox.

- Receiving Data: When a packet arrives, the server reads the data. It's like opening a letter and reading the message inside.

- Printing the Message: After reading the data, the server prints out the message. This is like sharing the letter's content with someone in the room.

- Continuous Listening: The server keeps doing this in a loop, always ready for the next packet.

## Create a UDP server

- go run udpserver.go

- test the UDP server
  - `echo -n "Hello UDP Server" | nc -u localhost 8080`
  - `nc -u localhost 5501` then send requests to the UDP server.
