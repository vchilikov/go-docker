package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintf(w, "Hello! Hostname: %s", hostname)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", index) // установим роутер
	fmt.Println(" http://localhost:9000")
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		panic(err)
	}
}
