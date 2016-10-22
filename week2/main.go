package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	requestString := r.URL.EscapedPath()[1:]
	requestURL, err := base64.StdEncoding.DecodeString(requestString)
	if err == nil {
		fmt.Println("Redirected to " + string(requestURL))
		http.Redirect(w, r, "http://"+string(requestURL), 302)
	} else {
		w.Write([]byte("Invalid base64 url"))
	}
}

func main() {
	fmt.Println("Listen on :8085")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err.Error())
	}
}
