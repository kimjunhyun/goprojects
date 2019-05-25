package main

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
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

// Weather is
type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		No          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		Pressure  float32 `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		SeaLevel  float32 `json:"sea_level"`
		GrndLevel float32 `json:"srnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  float32 `json:"dt"`
	Sys struct {
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise float32 `json:"sunrise"`
		Sunset  float32 `json:"sunset"`
	} `json:"sys"`
	No   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func main() {
	var sActive string
	sActive = "Inactive"

	Icnt := 0

	//날씨정보 얻기
	go GetWeather()
	// 딜레이
	time.Sleep(5 * time.Second)

	for {
		time.Sleep(1 * time.Second)

		//장비정보를 가져오는 루틴
		go DoJob(Icnt%4, sActive)

		if Icnt%4 == 0 {
			// 딜레이
			time.Sleep(5 * time.Second)
			sActive = DoJobBinary()

			fmt.Printf(sActive + strconv.Itoa(Icnt) + "\n")
			//sActive := onss
		}

		if Icnt > 100 {
			// 딜레이
			time.Sleep(5 * time.Second)
			//날씨정보 얻기
			go GetWeather()
			//Icnt = 0
		}
		fmt.Printf(sActive + "\n")
		// 딜레이
		time.Sleep(10 * time.Second)
		Icnt++
	}
}

//GetWeather 날씨정보를 가져오는 함수
func GetWeather() {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=seoul,KR&APPID=b9ee287555ef243090cd25c4b2e359ef&units=metric")

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(data))

	var szData = []byte(data)
	var weather Weather
	err = json.Unmarshal(szData, &weather)
	if err != nil {
		panic(err)
	}

	fmt.Printf("최고 기온: %.2f °C\n", weather.Main.TempMax)
	fmt.Printf("최저 기온: %.2f °C\n", weather.Main.TempMin)
	fmt.Printf("현재기온: %.2f °C\n", weather.Main.Temp)

	fmt.Printf("일출: %.2f °C\n", weather.Sys.Sunrise)
	fmt.Printf("일몰: %.2f °C\n", weather.Sys.Sunset)
	fmt.Printf("구름: %d \n", weather.Clouds.All)
	fmt.Printf("구름1: %s \n", weather.Weather[0].Main)
	fmt.Printf("구름2: %s \n", weather.Weather[0].Description)

	sockstr("if/" + weather.Weather[0].Main)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

//DoJob : analog 데이타를 가져오는
func DoJob(i int, sact string) {
	///
	///
	sNum := [4]string{"1", "2", "4", "5"}
	//av1 : 외기온도
	//av2 : 외기습도
	//av4 : 초미세먼지
	//av5 : 미세먼지

	//장비 주소 RESTful API 주소
	// urlStr := "https://223.171.51.167/api/rest/v1/protocols/bacnet/local/objects/analog-value/" + sNum[i] + "/properties/present-value"
	urlStr := "https://172.24.253.241/api/rest/v1/protocols/bacnet/local/objects/analog-value/" + sNum[i] + "/properties/present-value"

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}

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

	//log.Println(response.StatusCode)
	//log.Println(response.Status)
	//getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	// log.Println(string(data))
	//s := strings.Split(data, "\n")
	//log.Println(data[4])

	res := magok{}
	json.Unmarshal(data, &res)

	log.Println(res)
	log.Println(res.Value)
	log.Println(string(data))
	sock(res.Value, i, sact)
}

//DoJobBinary 포그머신 가동여부 확인
func DoJobBinary() string {
	///
	///
	//장비 주소 RESTful API 주소
	// urlStr := "https://223.171.51.167/api/rest/v1/protocols/bacnet/local/objects/binary-input/203/properties/present-value" //사육신 현장
	urlStr := "https://172.24.253.241/api/rest/v1/protocols/bacnet/local/objects/binary-input/203/properties/present-value" //마곡중앙광장 현장

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}

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

	//getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	res := magok{}
	json.Unmarshal(data, &res)

	log.Println(res)
	log.Println(res.Value)
	log.Println(string(data))

	return res.Value
}

func sock(str string, i int, sact string) {
	sName := [4]string{"pt", "ph", "pc", "pm"}

	conn, _err := net.Dial("tcp", "127.0.0.1:8081")
	text := sName[i] + "/" + str + "\n"
	if sName[i] == "pm" {
		if sact == "Inactive" {
			text = "bpm/" + str + "\n"
		} else {
			text = "apm/" + str + "\n"
		}
	}
	// send to socket
	fmt.Fprintf(conn, text+"\n")

	defer conn.Close()

	if _err != nil {
		log.Fatalln(_err)
	}

}

func sockstr(str string) {
	conn, _err := net.Dial("tcp", "127.0.0.1:8081")
	text := str + "\n"
	log.Println(text)
	// send to socket
	fmt.Fprintf(conn, text+"\n")

	defer conn.Close()

	if _err != nil {
		log.Fatalln(_err)
	}

}
