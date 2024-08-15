package main

import (
	"encoding/json"
	"io"
	"os"
)

// Сериализация структуры в JSON
// Описание: Напишите программу, которая сериализует структуру Person в
// формат JSON и выводит результат на экран.
// Методы или функции:
// func serializePerson(p Person) (string, error)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func serializePerson(p Person) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Десериализация JSON в структуру
// Описание: Напишите программу, которая десериализует JSON-строку в
// структуру Person и выводит результат на экран.
// Методы или функции:
// func deserializePerson(s string) (Person, error)

func deserializePerson(data string) (Person, error) {
	var p Person
	err := json.Unmarshal([]byte(data), &p)
	return p, err
}

// Работа с вложенными структурами
// Описание: Напишите программу, которая сериализует и десериализует
// структуру с вложенной структурой Address и выводит результат на экран.
// Методы или функции:
// func serializePerson(p Person) (string, error)
// func deserializePerson(data string) (Person, error)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person2 struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func serializePerson2(p Person2) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func deserializePerson2(data string) (Person, error) {
	var p Person
	err := json.Unmarshal([]byte(data), &p)
	return p, err
}

// Работа с массивами структур
// Описание: Напишите программу, которая сериализует и десериализует массив
// структур Person и выводит результат на экран.
// Методы или функции:
// func serializePeople(people []Person) (string, error)
// func deserializePeople(data string) ([]Person, error)

func serializePeople(people []Person) (string, error) {
	data, err := json.Marshal(people)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func deserializePeople(data string) ([]Person, error) {
	var people []Person
	err := json.Unmarshal([]byte(data), &people)
	return people, err
}

// Пропуск полей при сериализации
// Описание: Напишите программу, которая сериализует структуру Person,
// пропуская поля, которые имеют пустые значения.
// Методы или функции:
// func serializePerson(p Person) (string, error)

type Person3 struct {
	Name  string `json:"name"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

func serializePerson3(p Person3) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Обработка неизвестных полей
// Описание: Напишите программу, которая десериализует JSON-строку в структуру,
// игнорируя неизвестные поля.
// Методы или функции:
// func deserializePerson(data string) (Person, error)

type Person4 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func deserializePerson4(data string) (Person4, error) {
	var p Person4
	err := json.Unmarshal([]byte(data), &p)
	return p, err
}

// Использование JSON с картами
// Описание: Напишите программу, которая сериализует и десериализует карту
// map[string]interface{} в JSON и выводит результат на экран.
// Методы или функции:
// func serializeMap(data map[string]interface{}) (string, error)
// func deserializeMap(jsonStr string) (map[string]interface{}, error)

func serializeMap(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func deserializeMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}

// Чтение JSON из файла
// Описание: Напишите программу, которая читает JSON из файла, десериализует его
// в структуру Person и выводит результат на экран.
// Методы или функции:
// func deserializePersonFromFile(path string) (Person, error)

func readPersonFromFile(filename string) (Person, error) {
	var person Person
	file, err := os.Open(filename)
	if err != nil {
		return person, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return person, err
	}

	err = json.Unmarshal(data, &person)
	return person, err
}

// Запись JSON в файл
// Описание: Напишите программу, которая сериализует структуру Person в формат
// JSON и записывает результат в файл.
// Методы или функции:
// func serializePersonToFile(person Person, filename string) error
// func deserializePersonFromFile(filename string) (Person, error)

func writePersonToFile(filename string, p Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

// Изменение значений в JSON
// Описание: Напишите программу, которая читает JSON-данные из строки, изменяет
// значение определенного поля и выводит обновленный JSON на экран.
// Методы или функции:
// func updatePerson(data string, newAge int) (string, error)

func updatePersonAge(data string, newAge int) (string, error) {
	var person Person
	err := json.Unmarshal([]byte(data), &person)
	if err != nil {
		return "", err
	}

	person.Age = newAge

	updatedData, err := json.Marshal(person)
	if err != nil {
		return "", err
	}

	return string(updatedData), nil
}
