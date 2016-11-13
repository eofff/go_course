package main

import (
	"fmt"
	"net/http"

	"week4/controllers"
)

func main() {
	controller := controllers.IndexController{}
	http.HandleFunc("/", controller.IndexHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
