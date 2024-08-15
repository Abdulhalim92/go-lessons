package main

import "fmt"

func main() {
	textTasks()
	jsonTasks()
}

func textTasks() {
	// 1. Подсчет символов в файле
	count, err := countCharacters("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total characters: %d\n", count)

	// 2. Подсчет строк в файле
	lines, err := countLines("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total lines: %d\n", lines)

	// 3. Подсчет слов в файле
	text := "Hello world from Go!"
	wordCount := countWords(text)
	fmt.Printf("Total words: %d\n", wordCount)

	// 4. Запись строки в файл
	err = writeStringToFile("output.txt", "Hello, Go!")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("String written to file successfully")

	// 5. Чтение первой строки из файла
	line, err := readFirstLine("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("First line:", line)

	// 6. Копирование файлов
	err = copyFile("source.txt", "destination.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File copied successfully")

	// 7. Чтение и запись в файл
	err = readAndWriteToFile("user_input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Text written to file successfully")

	// 8. Обратное чтение содержимого файла
	content, err := reverseReadFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Reversed file content:\n", content)

	// 9. Конкатенация файлов
	err = concatenateFiles("file1.txt", "file2.txt", "output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Files concatenated successfully")

	// 10. Проверка существования файла
	if fileExists("example.txt") {
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}

	// 11. Чтение и подсчет уникальных слов в файле
	uniqueWords, err := countUniqueWords("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Unique words: %d\n", uniqueWords)

	// 12. Поиск и замена слова в файле
	err = replaceWordInFile("example.txt", "oldWord", "newWord")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Word replaced successfully")

	// 13. Подсчет частоты слов в файле
	frequencies, err := wordFrequency("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for word, freq := range frequencies {
		fmt.Printf("Word: %s, Frequency: %d\n", word, freq)
	}

	// 14. Переворачивание строк в файле
	err = reverseLines("example.txt", "output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Lines reversed and written to file successfully")

	// 15. Удаление пустых строк из файла
	err = removeEmptyLines("example.txt", "output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Empty lines removed successfully")

	// 16. Сравнение двух файлов на идентичность
	identical, err := compareFiles("file1.txt", "file2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if identical {
		fmt.Println("Files are identical")
	} else {
		fmt.Println("Files are different")
	}

	// 17. Чтение файла с конца
	err = readFileReverse("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 18. Подсчет количества строк с определенным словом
	count, err = countLinesWithWord("example.txt", "Go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Number of lines containing the word 'Go': %d\n", count)

	// 19. Генерация файла с повторяющимися строками
	err = generateRepeatedLinesFile("output.txt", "Hello, Go!", 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File generated successfully")

	// 20. Подсчет количества символов в каждой строке файла
	err = countCharactersPerLine("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 21. Поиск самого длинного слова в файле
	longestWord, err := findLongestWord("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Longest word: %s\n", longestWord)

	// 22. Подсчет частоты встречаемости символов в файле
	frequencies2, err := charFrequency("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for char, count := range frequencies2 {
		fmt.Printf("Character: %q, Frequency: %d\n", char, count)
	}

	// 23. Реверсирование слов в каждой строке файла
	err = reverseWordsInFile("example.txt", "output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Words reversed and written to file successfully")

	// 24. Подсчет количества слов в каждой строке файла
	err = countWordsPerLine("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 25. Найти и заменить слово в файле, с сохранением регистра
	err = replaceWordWithCase("example.txt", "oldWord", "newWord")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Word replaced successfully with case sensitivity")

	// 26. Поиск самого короткого слова в файле
	shortestWord, err := findShortestWord("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Shortest word: %s\n", shortestWord)

	// 27. Чтение и запись файла по частям
	err = copyFileInChunks("source.txt", "destination.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File copied in chunks successfully")

	// 28. Подсчет количества символов, слов и строк в файле
	chars, words, lines, err := countFileStats("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Characters: %d, Words: %d, Lines: %d\n", chars, words, lines)

	// 29. Поиск и удаление строки в файле
	err = deleteLineFromFile("example.txt", "line to delete")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Line deleted successfully")

	// 30. Сортировка строк в файле в алфавитном порядке
	err = sortFileLines("example.txt", "sorted_output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File lines sorted successfully")
}

func jsonTasks() {
	// 1. Сериализация структуры в JSON
	person := Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	jsonData, err := serializePerson(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", jsonData)

	// 2. Десериализация JSON в структуру
	jsonData = `{"name":"John Doe","age":30,"email":"john.doe@example.com"}`
	person, err = deserializePerson(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person)

	// 3. Работа с вложенными структурами
	person2 := Person2{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
		Address: Address{
			Street: "123 Main St",
			City:   "Springfield",
		},
	}
	jsonData2, err := serializePerson2(person2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", jsonData2)
	deserializedPerson2, err := deserializePerson2(jsonData2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deserialized Person: %+v\n", deserializedPerson2)

	// 4. Работа с массивами структур
	people := []Person{
		{Name: "John Doe", Age: 30, Email: "john.doe@example.com"},
		{Name: "Jane Smith", Age: 25, Email: "jane.smith@example.com"},
	}
	jsonData, err = serializePeople(people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", jsonData)
	deserializedPeople, err := deserializePeople(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deserialized People: %+v\n", deserializedPeople)

	// 5. Пропуск полей при сериализации
	person3 := Person3{Name: "John Doe", Age: 0, Email: ""}
	jsonData3, err := serializePerson3(person3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", jsonData3)

	// 6. Обработка неизвестных полей
	jsonData4 := `{"name":"John Doe","age":30,"email":"john.doe@example.com"}`
	person4, err := deserializePerson4(jsonData4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person4)

	// 7. Использование JSON с картами
	data := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john.doe@example.com",
	}
	jsonData, err = serializeMap(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", jsonData)
	deserializedMap, err := deserializeMap(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deserialized Map: %+v\n", deserializedMap)

	// 8. Чтение JSON из файла
	person, err = readPersonFromFile("person.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person)

	// 9. Запись JSON в файл
	person = Person{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	err = writePersonToFile("person.json", person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Person data written to file successfully")

	// 10. Изменение значений в JSON
	jsonData = `{"name":"John Doe","age":30,"email":"john.doe@example.com"}`
	updatedJson, err := updatePersonAge(jsonData, 35)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Updated JSON:", updatedJson)

}
