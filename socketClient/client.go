package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	// connect to this socket
	for {
		go sock("ssss")
		// listen for reply
		time.Sleep(10 * time.Second)
	}
}

func sock(str string) {
	conn, _err := net.Dial("tcp", "127.0.0.1:8081")
	text := str + "test\n"
	// send to socket
	fmt.Fprintf(conn, text+"\n")

	defer conn.Close()

	if _err != nil {
		log.Fatalln(_err)
	}

}
