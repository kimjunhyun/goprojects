package main

import (
	"net/http"

	"encoding/json"
	"log"

	"github.com/gorilla/mux"
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
