package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if err := client("127.0.0.1:9876"); err != nil {
		log.Fatalln("client error:", err)
	}
}

// a basic client implementation to connect to the tcp chat server
// returns an error if it cant connect to the server address or
// reading from stdin fails
func client(serverAddress string) error {
	// dial up the server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	addr := conn.LocalAddr()

	// print the local address as chat indicator
	fmt.Printf("%s: ", addr)

	// read incoming messages from the server without blocking
	go func(c net.Conn) {
		conScan := bufio.NewScanner(conn)
		for conScan.Scan() {
			// if a message is received, overwrite the previous line in the terminal with the message
			// and then print it. This ensures there is always a '<address>:' at the bottom of the terminal
			if _, err = fmt.Fprintf(os.Stdout, "\033[2K\r%s\n", conScan.Text()); err != nil {
				log.Println("error reading from server:", err)
			}
			fmt.Printf("%s: ", addr)
		}
		if err := conScan.Err(); err != nil {
			log.Println("connection scanner error:", err)

		}
	}(conn)

	// read all lines from stdin and
	// add a new address indicator after a message was written
	stdinScan := bufio.NewScanner(os.Stdin)
	for stdinScan.Scan() {
		if _, err := fmt.Fprintln(conn, stdinScan.Text()); err != nil {
			return err
		}
		fmt.Printf("%s: ", addr)
	}
	return stdinScan.Err()
}
