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

var cnt int

func health(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	cnt++
	if cnt > 4 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, "Hello! Hostname: %s Cnt: %d", hostname, cnt)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", health)
	fmt.Println(" http://localhost:9000")
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		panic(err)
	}
}
