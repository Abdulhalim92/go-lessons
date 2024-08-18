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

	//var address Address = Address{
	//	Street: "Rudaki",
	//	City:   "Dushanbe",
	//}
	//
	//var p Person2 = Person2{
	//	Name:    "Sobirjon",
	//	Age:     21,
	//	Email:   "sobirjon@go.dev",
	//	Address: address,
	//}
	//
	//data, err := json.MarshalIndent(p, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(data))
	//
	file, err := os.OpenFile("example.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	//
	//n, err := file.Write(data)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("Wrote %d bytes to file.\n", n)

	jsonData := `{
	"name": "Sobirjon",
	"age": 21,
	"email": "sobirjon@go.dev",
	"address": {
		"street": "Rudaki",
		"city": "Dushanbe"
		}
	}`

	byteData := []byte(jsonData)

	var person Person2
	err = json.Unmarshal(byteData, &person)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", person)

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var p2 Person2

	err = json.Unmarshal(data, &p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p2)

}

type Person2 struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
}

func rwToFile() {
	// Открываем файл для чтения и записи
	file, err := os.OpenFile("example.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var p2 []Person

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Если файл не пустой, десериализуем данные
	if len(data) > 0 {
		err = json.Unmarshal(data, &p2)
		if err != nil {
			panic(err)
		}
	}

	// Добавляем нового человека в массив
	p := Person{
		Name: "Said",
		Age:  12,
	}

	p2 = append(p2, p)

	// Возвращаем указатель на начало файла
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Сериализуем массив обратно в JSON
	data, err = json.MarshalIndent(p2, "", "    ")
	if err != nil {
		panic(err)
	}

	// Очищаем файл перед записью новых данных
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// Записываем данные обратно в файл
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")

}
