# Create a UDP server

- go run udpserver.go

- test the UDP server
  - `echo -n "Hello UDP Server" | nc -u localhost 8080`
  - `nc -u localhost 5501` then send requests to the UDP server.
