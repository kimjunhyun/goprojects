package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-zeromq/zmq4"
)

func main() {
	for {
		socketOpen("221.162.15.248:13687")
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func MQOpen() {
}

func socketOpen(sAddr string) {
	conn, err := net.Dial("tcp", sAddr)

	if nil != err {
		log.Println(err)
		return
	}

	data := make([]byte, 40960)
	bCon := make([]bool, 1)

	floor1 := make([]byte, 4096)
	floor2 := make([]byte, 4096)
	floor3 := make([]byte, 4096)
	floor4 := make([]byte, 4096)
	floor5 := make([]byte, 4096)
	floor6 := make([]byte, 4096)
	floor7 := make([]byte, 4096)
	floor8 := make([]byte, 4096)
	floor9 := make([]byte, 4096)

	bCon[0] = true

	go func() {
		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				bCon[0] = false

				return
			}

			parcingData(data, n, &floor1, &floor2, &floor3, &floor4, &floor5, &floor6, &floor7, &floor8, &floor9)
			//log.Println("Server send : " + string(data[:n]))
			//log.Println("Server send : " + string(n))
			str := string(floor2)
			fmt.Printf(str + "\n")

			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	ii := 0
	go func() {
		for {
			fmt.Printf("[+++1] \n")
			pub := zmq4.NewPub(context.Background())
			defer pub.Close()

			fmt.Printf("[+++2] \n")
			if ii == 0 {
				err := pub.Listen("tcp://127.0.0.1:5563")
				if err != nil {
					// log.Fatalf("could not listen: %v", err)
					fmt.Printf("could not listen: %v", err)
					//pub.Close()
					time.Sleep(3 * time.Second)
					continue
				}
			}
			ii = 1
			fmt.Printf("[+++3] \n")
			msgA := zmq4.NewMsgFrom(
				[]byte("1"), //This message has named "test1" key
				floor1,
			)
			// msgB := zmq4.NewMsgFrom(
			// 	[]byte("2"), //This message has named "test2" key
			// 	floor2,
			// )
			i := 0
			for {
				// err = pub.Send(msgA)
				i = i + 1
				fmt.Printf("[--] \n")
				// fmt.Printf("[%s] \n", floor1)
				msgA.Frames[1] = floor1 //[]byte("ss" + strconv.Itoa(i))
				err = pub.Send(msgA)
				if err != nil {
					// log.Fatal(err)
					break
				}
				time.Sleep(3 * time.Second)
			}
			fmt.Printf("[++] \n")
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()

	defer conn.Close()

	bByte := make([]byte, 9) //장비에게 전달할 배열
	for {

		if bCon[0] == false {
			return
		}

		{
			bByte[0] = 0x02
			bByte[1] = 0X0e
			bByte[2] = 0X30
			bByte[3] = 0X30
			bByte[4] = 0X52
			bByte[5] = 0X43
			bByte[6] = 0X4f
			bByte[7] = 0X4d
			bByte[8] = 0X03
		}
		conn.Write(bByte)

		time.Sleep(time.Duration(10) * time.Second)
	}
}

func parcingData(d []byte, iLen int, floor1 *[]byte, floor2 *[]byte, floor3 *[]byte, floor4 *[]byte, floor5 *[]byte, floor6 *[]byte, floor7 *[]byte, floor8 *[]byte, floor9 *[]byte) {
	//var buffer1 [4096]byte
	// buffer1 := make([]byte, 4096)
	var iType byte
	iType = '0'
	iCnt := 0
	for i := 0; i < iLen; i++ {
		//00A00DTU 01A01DTU
		if (d[i] == '0' && d[i+1] == '0' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '0') ||
			(d[i] == '0' && d[i+1] == '1' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '1') ||
			(d[i] == '0' && d[i+1] == '2' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '2') ||
			(d[i] == '0' && d[i+1] == '3' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '3') ||
			(d[i] == '0' && d[i+1] == '4' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '4') ||
			(d[i] == '0' && d[i+1] == '5' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '5') ||
			(d[i] == '0' && d[i+1] == '6' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '6') ||
			(d[i] == '0' && d[i+1] == '7' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '7') ||
			(d[i] == '0' && d[i+1] == '8' && d[i+2] == 'A' && d[i+3] == '0' && d[i+4] == '8') {
			// log.Println(string(d[i+1]) + "------------\n")
			iType = d[i+1]
			iCnt = 0
		}
		if d[i] == '0' && d[i+1] == '0' && d[i+2] == 'Z' && d[i+3] == 'S' && d[i+4] == 'T' && d[i+5] == 'A' && d[i+6] == 'R' && d[i+7] == 'T' {
			log.Println("Start\n")
		}
		if d[i] == 0x02 {
			iCnt = 0
		} else if d[i] == 0x03 {
			iCnt = 0
			//log.Println(str + "\n")
			//Array.Clear(buffer1, 0, buffer1.Length);
		} else {
			iCnt = iCnt + 1
			// buffer1[iCnt] = d[i]

			// (*floor1)[iCnt] = d[i]

			if iType == '0' {
				(*floor1)[iCnt] = d[i]
			} else if iType == '1' {
				(*floor2)[iCnt] = d[i]
			} else if iType == '2' {
				(*floor3)[iCnt] = d[i]
			} else if iType == '3' {
				(*floor4)[iCnt] = d[i]
			} else if iType == '4' {
				(*floor5)[iCnt] = d[i]
			} else if iType == '5' {
				(*floor6)[iCnt] = d[i]
			} else if iType == '6' {
				(*floor7)[iCnt] = d[i]
			} else if iType == '7' {
				(*floor8)[iCnt] = d[i]
			} else if iType == '8' {
				(*floor9)[iCnt] = d[i]
			}

		}

	}
	iType = '0'
}
