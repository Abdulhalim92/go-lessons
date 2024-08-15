package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	//in := bufio.NewReader(os.Stdin)
	//out := bufio.NewWriter(os.Stdout)
	//defer func() { _ = out.Flush() }()
	//
	//var num int
	//fmt.Fscan(in, &num)
	//
	//fmt.Fprintln(out, num)

	file, err := os.OpenFile("example.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var p2 []Person2

	err = json.Unmarshal(data, &p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", p2)

	p := Person2{
		Name:      "Said",
		Age:       12,
		IsStudent: true,
	}

	p2 = append(p2, p)

	data, err = json.MarshalIndent(p2, "", "    ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}

type Person2 struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsStudent bool   `json:"is_student"`
	Address   string `json:"address"`
}
