package main

import (
	"net/http"
	"encoding/json"
	"strings"
)

type request_json struct {
	Key string
}

type response_json struct {
	Count int `json:"count"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rj request_json
	decoder.Decode(&rj)
	a, _ := json.Marshal(response_json{len(strings.Split(rj.Key, " "))})
	w.Header().Set("Content-Type", "application/json")
	w.Write(a)
}

func main() {
	http.HandleFunc("/count", handler)
	http.ListenAndServe(":9000", nil)
}
