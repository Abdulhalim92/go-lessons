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

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, это мой HTTP-сервер на Go!\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func main() {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", handlerFunc)
	//mux.HandleFunc("/hello", sayHello)

	myHandler := NewMyHandler(mux)
	myHandler.mux.HandleFunc("/", handlerFunc)
	myHandler.mux.HandleFunc("/hello", sayHello)

	fmt.Printf("Server started: port: %s", ":8080\n")

	err := http.ListenAndServe(":8080", myHandler)
	if err != nil {
		panic(err)
	}

}

type MyHandler struct {
	mux *http.ServeMux
}

func NewMyHandler(mux *http.ServeMux) *MyHandler {
	return &MyHandler{mux: mux}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
