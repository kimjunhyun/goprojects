package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/jacobsa/go-serial/serial"

	gopher_and_rabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

//NewSerial 함수
func NewSerial(sport string, irate uint, idata uint, istop uint, chname string) {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        sport,
		BaudRate:        irate,
		DataBits:        idata,
		StopBits:        istop,
		MinimumReadSize: 4,
		ParityMode:      0,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	go requestHandler(port)
	// b := []byte("ABC€")
	// port.Write(b)
	// b = []byte(NewLine())
	// port.Write(b)
	// b = []byte(NewLine())
	// port.Write(b)
	// b = []byte(ConvertFontSize(6, 6))
	// port.Write(b)

	// GetKrString(" ")
	// port.Write([]byte(GetKrString(" 한글")))
	// time.Sleep(100 * time.Millisecond)
	// b = []byte(NewLine())
	// port.Write(b)
	// b = []byte(ConvertFontSize(1, 1))
	// port.Write(b)
	// port.Write([]byte(GetKrString("김준현")))
	// time.Sleep(100 * time.Millisecond)
	// b = []byte(NewLine())
	// port.Write(b)
	// port.Write(b)
	// port.Write(b)
	// port.Write(b)
	// port.Write(b)
	// port.Write(b)
	// port.Write(b)
	// time.Sleep(100 * time.Millisecond)
	// b = []byte(LineFeed(10))
	// port.Write(b)
	// time.Sleep(100 * time.Millisecond)
	// b = []byte(Cut())

	// Make sure to close it later.
	defer port.Close()

	// Write 4 bytes to the port.
	// b := []byte{0x00, 0x01, 0x02, 0x03}
	// n, err := port.Write(b)
	// if err != nil {
	// 	log.Fatalf("port.Write: %v", err)
	// }

	// fmt.Println("Wrote", n, "bytes.")
	// floor5 := make([]byte, 40)
	// floor5[0] = 'A'
	// fmt.Println("Wrote", floor5, "bytes.")

	fmt.Println("Wrote   bytes.")
	connQT, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	fmt.Println("Wrote   bytes.")
	handleError(err, "Can't connect to AMQP")
	defer connQT.Close()
	fmt.Println("Wrote   bytes.")

	amqpChannel, err := connQT.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare(chname, true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		//log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)
			if d.Body[0] == 'c' && d.Body[1] == 'm' && d.Body[2] == 'd' {
				sCommand := strings.Split(string(d.Body), ",")
				if sCommand[0] == "cmd" {
					go SerialWrite(port, sCommand[1])
				}
			}

			addTask := &gopher_and_rabbit.AddTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

			body := "Hello World!"
			err = amqpChannel.Publish(
				"",     // exchange
				chname, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
		}
	}()

	// Stop for program termination
	<-stopChan
}

func requestHandler(c io.ReadWriteCloser) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		// _, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

	}
}

//SerialWrite 함수
func SerialWrite(c io.ReadWriteCloser, str string) {

	_, err := c.Write([]byte(str)) // 클라이언트로 데이터를 보냄
	if err != nil {
		fmt.Println(err)
		return
	}
}
