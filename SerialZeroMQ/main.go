/*
 * This application can be used to experiment and test various serial port options
 */

package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"time"

	"golang.org/x/text/encoding/korean"

	"golang.org/x/text/transform"

	"github.com/jacobsa/go-serial/serial"

)

func main() {
	strPort := flag.String("port", "COM5", "a string")
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
	b := []byte("ABC€" + msgA.String())
	port.Write(b)
	b = []byte(NewLine())
	port.Write(b)
	b = []byte(NewLine())
	port.Write(b)
	b = []byte(ConvertFontSize(6, 6))
	port.Write(b)

	GetKrString(" ")
	port.Write([]byte(GetKrString(" 한글")))
	time.Sleep(100 * time.Millisecond)
	b = []byte(NewLine())
	port.Write(b)
	b = []byte(ConvertFontSize(1, 1))
	port.Write(b)
	port.Write([]byte(GetKrString("김준현")))
	time.Sleep(100 * time.Millisecond)
	b = []byte(NewLine())
	port.Write(b)
	port.Write(b)
	port.Write(b)
	port.Write(b)
	port.Write(b)
	port.Write(b)
	port.Write(b)
	time.Sleep(100 * time.Millisecond)
	b = []byte(LineFeed(10))
	port.Write(b)
	time.Sleep(100 * time.Millisecond)
	b = []byte(Cut())
	port.Write(b)

	// Make sure to close it later.
	defer port.Close()

}

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
