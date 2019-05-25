/*
 * This application can be used to experiment and test various serial port options
 */

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"

	"golang.org/x/text/encoding/korean"

	"golang.org/x/text/transform"

	"github.com/jacobsa/go-serial/serial"

	gopher_and_rabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func main() {

	//go SetRabbit()

	strPort := flag.String("port", "COM3", "a string")
	numBaudRate := flag.Int("baudrate", 9600, "an int")
	numDataBits := flag.Int("DataBits", 8, "an int")
	numStopBits := flag.Int("StopBits", 1, "an int")
	flag.Parse()

	// Set up options.
	options := serial.OpenOptions{
		PortName:        *strPort,
		BaudRate:        uint(*numBaudRate),
		DataBits:        uint(*numDataBits),
		StopBits:        uint(*numStopBits),
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

	queue, err := amqpChannel.QueueDeclare("안녕하세요", true, false, false, false, nil)
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
			//log.Printf("Received a message: %s", d.Body)
			if string(d.Body) == "received" {
				go SerialWrite(port, "ffff ")
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
				"",      // exchange
				"안녕하세요", // routing key
				false,   // mandatory
				false,   // immediate
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

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func SerialWrite(c io.ReadWriteCloser, str string) {

	_, err := c.Write([]byte(str)) // 클라이언트로 데이터를 보냄
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func SetRabbit() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	handleError(err, "Can't connect to AMQP")
// 	defer conn.Close()

// 	amqpChannel, err := conn.Channel()
// 	handleError(err, "Can't create a amqpChannel")

// 	defer amqpChannel.Close()

// 	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
// 	handleError(err, "Could not declare `add` queue")

// 	rand.Seed(time.Now().UnixNano())

// 	addTask := gopher_and_rabbit.AddTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
// 	body, err := json.Marshal(addTask)
// 	if err != nil {
// 		handleError(err, "Error encoding JSON")
// 	}

// 	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
// 		DeliveryMode: amqp.Persistent,
// 		ContentType:  "text/plain",
// 		Body:         body,
// 	})

// 	handleError(err, "Error publishing message: %s")

// 	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)

// }

func ConvertFontSize(width int, height int) string {
	result := "0"
	_w := 0
	_h := 0

	//가로변환
	if width == 1 {
		_w = 0
	} else if width == 2 {
		_w = 16
	} else if width == 3 {
		_w = 32
	} else if width == 4 {
		_w = 48
	} else if width == 5 {
		_w = 64
	} else if width == 6 {
		_w = 80
	} else if width == 7 {
		_w = 96
	} else if width == 8 {
		_w = 112
	} else {
		_w = 0
	}
	//세로변환
	if height == 1 {
		_h = 0
	} else if height == 2 {
		_h = 1
	} else if height == 3 {
		_h = 2
	} else if height == 4 {
		_h = 3
	} else if height == 5 {
		_h = 4
	} else if height == 6 {
		_h = 5
	} else if height == 7 {
		_h = 6
	} else if height == 8 {
		_h = 7
	} else {
		_h = 0
	}
	//가로x세로
	sum := _w + _h
	result = string(29) + "!" + string(sum)

	return result
}

func GetKrString(s string) string {
	var bufs bytes.Buffer

	bufs.Reset()

	wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())

	wr.Write([]byte(s))

	wr.Close()

	return bufs.String()
}
func LineFeed(d int) string {
	strData := string(27) + "b" + string(d)
	return strData
}
func Cut() string {
	strData := string(29) + "V" + string(1)
	return strData
}
func NewLine() string {
	strData := string(10)
	return strData
}
