package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-zeromq/zmq4"
)

func main() {
	log.SetPrefix("psenvsub: ")
	//  Prepare our subscriber
	sub := zmq4.NewSub(context.Background())
	defer sub.Close()
	err := sub.Dial("tcp://localhost:5563")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	err = sub.SetOption(zmq4.OptionSubscribe, "1") //Received only named "test1" key.
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}
	for {
		// Read envelope
		msg, err := sub.Recv()
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		//log.Printf("[%s] %s\n", msg.Frames[0], msg.Frames[1])
		fmt.Printf("[%s] %s\n", msg.Frames[0], msg.Frames[1])

		//sub.Send()
	}
}
