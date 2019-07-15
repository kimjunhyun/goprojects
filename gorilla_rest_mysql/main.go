package main

import (
	"fmt"
	"net/http"
	"os"

	"encoding/json"
	"log"

	"github.com/gorilla/mux"

	//ini 사용을 위한
	"github.com/go-ini/ini"
)

type data struct {
	code        string
	Title       string
	Description string
}

var testData []data

func main() {
	testData = append(testData, data{code: "1", Title: "fist title", Description: "desc"})
	testData = append(testData, data{code: "2", Title: "second title", Description: "desc"})
	router := mux.NewRouter()
	router.HandleFunc("/test/{code}", GetData).Methods("Get")
	http.ListenAndServe(":8080", httpHandler(router))

	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	host := cfg.Section("db").Key("host").String()
	database := cfg.Section("db").Key("database").String()
	user := cfg.Section("db").Key("user").String()
	password := cfg.Section("db").Key("password").String()

	CreateDB(user, password, host, database)
}

func httpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}

func GetData(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	for _, i := range testData {
		if i.code == p["code"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	//json.NewEncoder(w).Encode(&event{})
}
