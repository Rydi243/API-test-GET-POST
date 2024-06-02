package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello world")
}

func kv_ch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kv_ch" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	numberStr := r.FormValue("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	square := number * number
	fmt.Fprintf(w, "Квадрат числа %d это %d", number, square)
}

func ch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ch" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}

	maps := make(map[int]string)

	chStr := r.FormValue("ch")
	ch, err := strconv.Atoi((chStr))
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	for i := 1; i <= ch; i++ {
		maps[i] = "манул"
		fmt.Fprintf(w, "%d вышел %s\n", i, maps[i])
	}

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/kv_ch", kv_ch)
	http.HandleFunc("/ch", ch)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
