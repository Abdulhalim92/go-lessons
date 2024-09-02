package main

import (
	"bytes"
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

func main() {
	client := http.Client{}

	person := Person{
		Name:    "Sasha",
		Age:     21,
		Address: "Moscow",
	}

	data, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(data)

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", buf)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)

	data, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
