package main

import (
	"database/sql"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
const (
	host     = "127.0.0.1"
	database = "test"
	user     = "root"
	password = "admin"
)

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
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=bucheon,KR&APPID=b9ee287555ef243090cd25c4b2e359ef&units=metric")

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

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
func db(msg string) {

	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")

	// Insert some data into table.
	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("banana", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("orange", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	if len(msg) < 5 {
		res, err = sqlStatement.Exec("apple", msg)
		checkError(err)
		rowCount, err = res.RowsAffected()
		fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
		fmt.Println("Done.")
	}
