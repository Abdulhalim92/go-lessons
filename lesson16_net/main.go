package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if len(data) == 0 {
		fmt.Println("Ничего не получено")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ничего не получено"))
		return
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", handlerFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server started: port: %s", ":8080")
}
