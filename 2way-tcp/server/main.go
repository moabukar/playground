package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	if err := listenAndServeTCP(context.TODO(), "127.0.0.1:9876"); err != nil {
		log.Fatalln("server error:", err)
	}
}

// runs a tcp server and returns an error if it could not bind the address
func listenAndServeTCP(ctx context.Context, address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// create a new message broker and run it concurrently
	b := newBroker(listener.Addr())
	go func() { b.run(ctx) }()

	// create a new connection handler that is used
	// for incoming tcp connections
	connectionHandler := newConnectionHandler(b)

	// accept all incoming connection in a loop
	// and dispatch the to the connection handler
	// as separate goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}
		go connectionHandler(conn)
	}
}

// the message is used to broadcast messages
// from a client to all other clients
type message struct {
	text   string
	sender net.Addr
}

func (m *message) write(w io.Writer) (n int, err error) {
	return fmt.Fprintf(w, "%s: %s\n", m.sender, m.text)
}

// the subscription is used to (un)subscribe to the broker
type subscription struct {
	msgChan chan *message
	addr    net.Addr
}

// the broker implements a fanout pattern
// it sends messages from the broadcast channel
// to the message channels in the clients map.
// the sub and unsub channels are used to add/remove
// a clients message channel to the list of clients
// while avoiding data races. This avoids the need for locks
type broker struct {
	serverAddr    net.Addr
	clients       map[chan *message]net.Addr
	subChan       chan *subscription
	unsubChan     chan *subscription
	broadcastChan chan *message
}

func newBroker(serverAddr net.Addr) *broker {
	return &broker{
		serverAddr:    serverAddr,
		clients:       make(map[chan *message]net.Addr, 10),
		subChan:       make(chan *subscription),
		unsubChan:     make(chan *subscription),
		broadcastChan: make(chan *message),
	}
}

// select from the control channels until the context is done
func (b *broker) run(ctx context.Context) error {
	for {
		select {
		// if the context is done, return
		case <-ctx.Done():
			return ctx.Err()

		// if a subscription is received from the subscribe channel,
		// add it to the clients map and send a welcome message
		case s := <-b.subChan:
			log.Printf("%s subscribed\n", s.addr)
			b.clients[s.msgChan] = s.addr
			s.msgChan <- &message{"welcome", b.serverAddr}

		// if a subscription is received from the unsubscribe channel,
		// close its message channel and delete if from the clients map
		case s := <-b.unsubChan:
			log.Printf("%s: is gone\n", s.addr)
			close(s.msgChan)
			delete(b.clients, s.msgChan)

		// if a message is received from the broadcast channel,
		// sent it to all  message channels in the clients map
		case msg := <-b.broadcastChan:
			log.Printf("%s broadcasting msg", msg.sender)
			for c := range b.clients {
				c <- msg
			}
		}
	}
}

// return a new connection handler to handle a tcp connection
func newConnectionHandler(b *broker) func(net.Conn) {
	return func(conn net.Conn) {
		// make sure the connection is released
		// when this function returns
		defer conn.Close()

		// read the address only once
		addr := conn.RemoteAddr()

		// create a new subscription to
		// (un)subscribe and receive messages
		msgChan := make(chan *message, 1)
		sub := &subscription{msgChan, addr}

		// subscribe to the broker to receive broadcast messages
		b.subChan <- sub

		// make sure to unsubscribe if the function returns
		// the function will return if a client disconnects
		defer func() { b.unsubChan <- sub }()

		// send all messages from the message channel to the connection
		// messages received on this channel have been broadcasted by
		// one of the connected clients
		go func() {
			for msg := range msgChan {
				// don't send the messages from this client to the client
				// the client already know what they sent
				if msg.sender == addr {
					continue
				}
				// otherwise sent the message
				if _, err := msg.write(conn); err != nil {
					log.Println("error sending message to client:", err)
				}
			}
		}()

		// read all messages from the connection and broadcast them
		// the message will be sent to all connected clients
		s := bufio.NewScanner(conn)
		for s.Scan() {
			b.broadcastChan <- &message{s.Text(), addr}
		}
		if err := s.Err(); err != nil {
			log.Println("error scanning next line:", err)
		}
	}
}
