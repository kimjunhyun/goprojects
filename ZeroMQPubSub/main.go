// main.go
package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-zeromq/zmq4"
)

func main() {
	// prepare the publisher
	pub := zmq4.NewPub(context.Background())
	defer pub.Close()

	err := pub.Listen("tcp://127.0.0.1:5563")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	msgA := zmq4.NewMsgFrom(
		[]byte("test1"), //This message has named "test1" key
		[]byte("My key is test1"),
	)
	msgB := zmq4.NewMsgFrom(
		[]byte("test2"), //This message has named "test2" key
		[]byte("my key is test2"),
	)
	i := 0
	for {
		// err = pub.Send(msgA)
		i = i + 1
		msgA.Frames[1] = []byte("ss" + strconv.Itoa(i))
		err = pub.Send(msgA)
		if err != nil {
			log.Fatal(err)
		}
		err = pub.Send(msgB)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(3 * time.Second)
	}
}
