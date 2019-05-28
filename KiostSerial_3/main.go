/*
 * This application can be used to experiment and test various serial port options
 */

package main

import (
	"bytes"
	"flag"
	"log"

	"golang.org/x/text/encoding/korean"

	"golang.org/x/text/transform"
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

	NewSerial(*strPort, uint(*numBaudRate), uint(*numDataBits), uint(*numStopBits))
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
		_w = 0 //sdklhskjgh
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
