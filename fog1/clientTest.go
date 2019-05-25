package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//마곡 장비 정보를 받기 위한 구조체
type magok struct {
	Value          string
	AsnValue       string
	IsArray        bool
	Implementation string
	Name           string
}

func main() {
	icnt := 0
	for {
		//장비정보를 가져오는 루틴
		go DoJob(icnt % 4)
		// 딜레이
		time.Sleep(10 * time.Second)
		icnt++
		if icnt > 100 {
			//날씨정보 얻기

			time.Sleep(10 * time.Second)
			icnt = 0
		}
	}
}

func DoJob(i int) {
	///
	///
	//sNum := [4]string{"1", "2", "4", "5"}
	//av1 : 외기온도
	//av2 : 외기습도
	//av4 : 초미세먼지
	//av5 : 미세먼지

	//장비 주소 RESTful API 주소
	// urlStr := "https://223.171.51.167/api/rest/v1/protocols/bacnet/local/objects/analog-value/" + sNum[i] + "/properties/present-value"
	urlStr := "https://223.171.51.167/api/rest/v1/protocols/bacnet/local/objects/binary-input/203/properties/present-value"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}
	//ai,av ,bi,bv1
	//adding the proxy settings to the Transport object
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }

	//adding the Transport object to the http Client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//	client := &http.Client{}

	//generating the HTTP GET request
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}
	//1
	//
	//adding proxy authentication
	auth := "user:Eclypse1234"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Authorization", basicAuth)

	//printing the request to the console
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Status)
	//getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	log.Println(string(data))
}

func sock(str string, i int) {
	sName := [4]string{"pt", "ph", "pc", "pm"}

	conn, _err := net.Dial("tcp", "127.0.0.1:8081")
	text := sName[i] + "/" + str + "\n"
	// send to socket
	fmt.Fprintf(conn, text+"\n")

	defer conn.Close()

	if _err != nil {
		log.Fatalln(_err)
	}

}
