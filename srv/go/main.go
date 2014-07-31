package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type request_json struct {
	Key string
}

type response_json struct {
	Count int `json:"count"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var rj request_json
	json.Unmarshal(body, &rj)
	a, _ := json.Marshal(response_json{len(strings.Split(rj.Key, " "))})
	fmt.Fprintf(w, string(a))
}

func main() {
	http.HandleFunc("/count", handler)
	http.ListenAndServe(":9000", nil)
}
