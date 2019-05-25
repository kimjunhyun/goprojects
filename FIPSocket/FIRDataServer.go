package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//go ServerSckOpen()
	for {
		socketOpen("221.162.15.248:13687")
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func requestHandler(c net.Conn, floor1 *[]byte, floor2 *[]byte, floor3 *[]byte, floor4 *[]byte, floor5 *[]byte, floor6 *[]byte, floor7 *[]byte, floor8 *[]byte, floor9 *[]byte) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		str := ""
		switch string(data[:n]) {
		case "floor1":
			str = string(*floor1)
		case "floor2":
			str = string(*floor2)
		case "floor3":
			str = string(*floor3)
		case "floor4":
			str = string(*floor4)
		case "floor5":
			str = string(*floor5)
		case "floor6":
			str = string(*floor6)
		case "floor7":
			str = string(*floor7)
		case "floor8":
			str = string(*floor8)
		case "floor9":
			str = string(*floor9)
		default:
			str = ""
		}

		_, err = c.Write([]byte(str)) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}

		//db(string(data[:n])) //데이타를 저장함

	}
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

			// str := string(floor5)
			// fmt.Printf(str + "\n")

			time.Sleep(time.Duration(10) * time.Second)
		}
	}()

	defer conn.Close()

	go func() {
		for {
			ln, err := net.Listen("tcp", ":1444") // TCP 프로토콜에 8000 포트로 연결을 받음
			if err != nil {
				fmt.Println(err)
				return
			}
			defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음

			for {
				conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

				go requestHandler(conn, &floor1, &floor2, &floor3, &floor4, &floor5, &floor6, &floor7, &floor8, &floor9) // 패킷을 처리할 함수를 고루틴으로 실행
			}

		}
	}()

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

		time.Sleep(time.Duration(60) * time.Second)
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
