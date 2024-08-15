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
