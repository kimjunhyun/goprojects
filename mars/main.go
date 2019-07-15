package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
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

var aa string

// Weather is
type Weather struct {
	Wea struct {
		Temp struct {
			Av float32 `json:"av"`
			Ct float32 `json:"ct"`
			Mn float32 `json:"mn"`
			Mx float32 `json:"mx"`
		} `json:"AT"`
		FUtc      string `json:"First_UTC"`
		WindSpeed struct {
			Av float32 `json:"av"`
			Ct float32 `json:"ct"`
			Mn float32 `json:"mn"`
			Mx float32 `json:"mx"`
		} `json:"WHS"`
		LUtc string `json:"Last_UTC"`
		Pre  struct {
			Av float32 `json:"av"`
			Ct float32 `json:"ct"`
			Mn float32 `json:"mn"`
			Mx float32 `json:"mx"`
		} `json:"PRE"`
		Season string `json:"Season"`
		Wd     struct {
			D0 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"0"`
			D1 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"1"`
			D10 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"10"`
			D11 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"11"`
			D12 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"12"`
			D13 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"13"`
			D14 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"14"`
			D15 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"15"`
			D2 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"2"`
			D3 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"3"`
			D5 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"5"`
			D6 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"6"`
			D7 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"7"`
			D8 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"8"`
			D9 struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"9"`
			Dmost_common struct {
				Cd float32 `json:"compass_degrees"`
				Cp string  `json:"compass_point"`
				Cr float32 `json:"compass_right"`
				Cu float32 `json:"compass_up"`
				Ct float32 `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	}
	Sol_Keys []string `json:"sol_keys"`
}

type Weather1 struct {
	Temp struct {
		Av float32
		Ct float32
		Mn float32
		Mx float32
	}
	FUtc      string
	WindSpeed struct {
		Av float32
		Ct float32
		Mn float32
		Mx float32
	}
	LUtc string
	Pre  struct {
		Av float32
		Ct float32
		Mn float32
		Mx float32
	}
	Season string
	Wd     struct {
		D0 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct int
		}
		D1 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D10 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D11 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D12 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D13 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D14 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D15 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D2 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D3 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D5 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D6 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D7 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D8 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		D9 struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
		Dmost_common struct {
			Cd float32
			Cp string
			Cr float32
			Cu float32
			Ct float32
		}
	}
}

type Vd struct {
	ahwd  [24]int
	valid bool
}
type Dd struct {
	At  Vd
	Hws Vd
	Pre Vd
	Wd  Vd
}
type MarsWD struct {
	Day1     Weather1
	Day2     Weather1
	Day3     Weather1
	Day4     Weather1
	Day5     Weather1
	Day6     Weather1
	Day7     Weather1 `json:"178"`
	Sol_Keys []string
	As       struct {
		DDay1 Dd
		DDay2 Dd
		DDay3 Dd
		DDay4 Dd
		DDay5 Dd
		DDay6 Dd
		DDay7 Dd
		DDay8 Dd
		Sol   int
		Sols  []string
	}
}

func main() {
	for {
		//장비정보를 가져오는 루틴
		go DoJob()
		// 딜레이
		time.Sleep(30 * time.Second)
	}
}

func DoJob() {
	///
	///
	//sNum := [4]string{"1", "2", "4", "5"}
	//av1 : 외기온도
	//av2 : 외기습도
	//av4 : 초미세먼지
	//av5 : 미세먼지

	//장비 주소 RESTful API 주소
	// urlStr := "https://223.171.51.167/api/rest/v1/protocols/bacnet/local/objects/analog-value/" + sNum[i] + "/properties/present-value"
	urlStr := "https://mars.nasa.gov/rss/api/?feed=weather&category=insight&feedtype=json&ver=1.0"
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
	//auth := "user:Eclypse1234"
	// basicAuth := "Basic " + "qp^2xo_(kiwv)+t500ljrhazx(wvt)16zbek1#e4tnaw_i%c85" //base64.StdEncoding.EncodeToString([]byte(auth))
	// request.Header.Add("Authorization", basicAuth)

	// //printing the request to the console
	// dump, _ := httputil.DumpRequest(request, false)
	// fmt.Println(string(dump))

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

	res1 := Weather{}
	json.Unmarshal([]byte(data), &res1)

	var res map[string]interface{}
	// var res MarsWD
	json.Unmarshal([]byte(data), &res)

	// log.Println(res["178.AT"])

	//log.Println(res.As.Sol)

	// log.Println(res)

	var sDataTot string

	sDataTot = ""

	log.Println("SOL = ", res1.Sol_Keys[6])
	sDataTot += res1.Sol_Keys[6]

	rD := res[res1.Sol_Keys[6]].(map[string]interface{})["Season"] //First_UTC
	str := fmt.Sprintf("%v", rD)
	log.Println("계절 = ", str)
	sDataTot += "," + str

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["AT"].(map[string]interface{})["av"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ := strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("ave Temp = ", fdata, "°C")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["AT"].(map[string]interface{})["mn"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("min Temp = ", fdata, "°C")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["AT"].(map[string]interface{})["mx"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("max Temp = ", fdata, "°C")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["HWS"].(map[string]interface{})["av"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("min Wind Speed = ", fdata, "m/s")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["HWS"].(map[string]interface{})["mn"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("min Wind Speed = ", fdata, "m/s")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["HWS"].(map[string]interface{})["mx"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("max Wind Speed = ", fdata, "m/s")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["PRE"].(map[string]interface{})["av"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("min Pressure = ", fdata/, "Pa")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["PRE"].(map[string]interface{})["mn"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("min Pressure = ", fdata, "Pa")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["PRE"].(map[string]interface{})["mx"]
	str = fmt.Sprintf("%v", rD)
	fdata, _ = strconv.ParseFloat(str, 64)
	// fdata = math.Round(fdata)
	log.Println("max Pressure = ", fdata, "Pa")
	sDataTot += "," + fmt.Sprintf("%.1f", fdata)

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["WD"].(map[string]interface{})["most_common"].(map[string]interface{})["compass_point"]
	str = fmt.Sprintf("%v", rD)
	log.Println("Wind Direction = ", str)
	sDataTot += "," + str

	rD = res[res1.Sol_Keys[6]].(map[string]interface{})["First_UTC"] //First_UTC
	str = fmt.Sprintf("%v", rD)
	log.Println("날짜 = ", str)
	sDataTot += "," + str

	log.Println(sDataTot)
	sockstr(sDataTot)
	// dd, err := strconv.ParseFloat(rD, 64)
	// log.Println(math.Round(rD))
	//log.Println(string(data))

	// log.Println(res.Sol_Keys)
	// log.Println(res.Wea.Temp.Av)
	// log.Println(res[0].Temp.Av)

	// log.Println(string(data))

	//printing the response
	//log.Println(string(data))
}

func sockstr(str string) {
	conn, _err := net.Dial("tcp", "127.0.0.1:14400")
	text := str + ""
	log.Println(text)
	// send to socket
	fmt.Fprintf(conn, text+"")

	defer conn.Close()

	if _err != nil {
		log.Fatalln(_err)
	}

}
