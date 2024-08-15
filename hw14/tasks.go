package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

// Подсчет символов в файле
// Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
// Методы или функции:
// func countCharacters(fileName string) (int, error)

func countCharacters(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	buf := make([]byte, 1024)
	totalCount := 0

	for {
		n, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return 0, err
		}
		totalCount += n
		if n == 0 {
			break
		}
	}

	return totalCount, nil
}

// Подсчет строк в файле
// Описание: Напишите программу, которая считает количество строк в текстовом файле.
// Методы или функции:
// func countLines(fileName string) (int, error)

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

// Счетчик слов в строке
// Описание: Напишите функцию, которая считает количество слов в строке.
// Методы или функции:
// func countWords(s string) int

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

//  Запись строки в файл
// Описание: Напишите программу, которая записывает заданную строку в файл.
// Методы или функции:
// func writeStringToFile(fileName, content string) error

func writeStringToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// Чтение первой строки файла
// Описание: Напишите программу, которая читает первую строку из текстового файла.
// Методы или функции:
// func readFirstLine(fileName string) (string, error)

func readFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

// Копирование содержимого одного файла в другой
// Описание: Напишите программу, которая копирует содержимое одного файла в другой.
// Методы или функции:
// func copyFile(src, dst string) error

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

//  Чтение строки с консоли и запись в файл
// Описание: Напишите программу, которая читает строку с консоли и записывает её в файл.
// Методы или функции:
// func readAndWriteToFile(fileName string) error

func readAndWriteToFile(fileName string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

// Обратное чтение файла
// Описание: Напишите программу, которая читает файл с конца в начало и выводит его содержимое на экран.
// Методы или функции:
// func reverseReadFile(fileName string) (string, error)

func reverseReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	n, err := file.ReadAt(buffer, 0)
	if err != nil && err != io.EOF {
		return "", err
	}

	reverseBuffer := make([]byte, n)
	for i := 0; i < n; i++ {
		reverseBuffer[i] = buffer[n-1-i]
	}

	return string(reverseBuffer), nil
}

// Объединение содержимого двух файлов
// Описание: Напишите программу, которая объединяет содержимое двух файлов и записывает результат в новый файл.
// Методы или функции:
// func concatenateFiles(file1, file2, outputFile string) error

func concatenateFiles(file1, file2, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, fileName := range []string{file1, file2} {
		inFile, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer inFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
	}

	return nil
}

// Проверка существования файла
// Описание: Напишите функцию, которая проверяет, существует ли файл с заданным именем.
// Методы или функции:
// func fileExists(fileName string) bool

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// Чтение и подсчет уникальных слов в файле
// Описание: Напишите программу, которая читает файл и подсчитывает количество уникальных слов.
// Методы или функции:
// func countUniqueWords(fileName string) (int, error)

func countUniqueWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	wordSet := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordSet[word] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return len(wordSet), nil
}

// Поиск и замена слова в файле
// Описание: Напишите программу, которая заменяет все вхождения определенного слова в файле на другое слово.
// Методы или функции:
// func replaceWordInFile(fileName, oldWord, newWord string) error

func replaceWordInFile(fileName, oldWord, newWord string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	updatedContent := strings.ReplaceAll(string(content), oldWord, newWord)

	return os.WriteFile(fileName, []byte(updatedContent), 0644)
}

// Подсчет частоты слов в файле
// Описание: Напишите программу, которая подсчитывает частоту каждого слова в файле.
// Методы или функции:
// func wordFrequency(fileName string) (map[string]int, error)

func wordFrequency(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordFreq[word]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}

// Переворачивание строк в файле
// Описание: Напишите программу, которая переворачивает строки в файле и записывает их в новый файл.
// Методы или функции:
// func reverseLines(fileName, outputFile string) error

func reverseLines(fileName, outputFile string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()
		reversedLine := reverseString(line)
		_, err = writer.WriteString(reversedLine + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Удаление пустых строк из файла
// Описание: Напишите программу, которая удаляет все пустые строки из файла.
// Методы или функции:
// func removeEmptyLines(fileName, outputFile string) error

func removeEmptyLines(fileName, outputFile string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Сравнение двух файлов на идентичность
// Описание: Напишите программу, которая сравнивает два файла и определяет, идентичны ли они.
// Методы или функции:
// func compareFiles(file1, file2 string) (bool, error)

func compareFiles(file1, file2 string) (bool, error) {
	content1, err := os.ReadFile(file1)
	if err != nil {
		return false, err
	}

	content2, err := os.ReadFile(file2)
	if err != nil {
		return false, err
	}

	return bytes.Equal(content1, content2), nil
}

// Чтение файла с конца
// Описание: Напишите программу, которая читает файл с конца и выводит его содержимое на экран.
// Методы или функции:
// func readFileReverse(fileName string) error

func readFileReverse(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, 1)

	for i := fileSize - 1; i >= 0; i-- {
		_, err := file.ReadAt(buffer, i)
		if err != nil {
			return err
		}
		fmt.Print(string(buffer))
	}

	return nil
}

// Подсчет количества строк с определенным словом
// Описание: Напишите программу, которая считает количество строк в файле, содержащих определенное слово.
// Методы или функции:
// func countLinesWithWord(fileName, word string) (int, error)

func countLinesWithWord(fileName, word string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			lineCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

// Генерация файла с повторяющимися строками
// Описание: Напишите программу, которая генерирует файл, содержащий заданную строку, повторенную указанное количество раз.
// Методы или функции:
// func generateRepeatedLinesFile(fileName, line string, count int) error

func generateRepeatedLinesFile(fileName, line string, count int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Подсчет количества символов в каждой строке файла
// Описание: Напишите программу, которая подсчитывает количество символов в каждой строке файла и выводит их на экран.
// Методы или функции:
// func countCharactersPerLine(fileName string) error

func countCharactersPerLine(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %d characters\n", lineNumber, len(line))
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Поиск самого длинного слова в файле
// Описание: Напишите программу, которая находит самое длинное слово в файле и выводит его на экран.
// Методы или функции:
// func findLongestWord(fileName string) (string, error)

func findLongestWord(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var longestWord string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return longestWord, nil
}

// Подсчет частоты встречаемости символов в файле
// Описание: Напишите программу, которая подсчитывает частоту встречаемости каждого символа в файле.
// Методы или функции:
// func charFrequency(fileName string) (map[rune]int, error)

func charFrequency(fileName string) (map[rune]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	frequency := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := []rune(scanner.Text())[0]
		frequency[char]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return frequency, nil
}

// Реверсирование слов в каждой строке файла
// Описание: Напишите программу, которая реверсирует слова в каждой строке файла и записывает результат в новый файл.
// Методы или функции:
// func reverseWordsInFile(fileName, outputFile string) error

func reverseWordsInFile(fileName, outputFile string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		_, err := writer.WriteString(strings.Join(words, " ") + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Подсчет количества слов в каждой строке файла
// Описание: Напишите программу, которая подсчитывает количество слов в каждой строке файла и выводит результат на экран.
// Методы или функции:
// func countWordsPerLine(fileName string) error

func countWordsPerLine(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		wordCount := len(strings.Fields(line))
		fmt.Printf("Line %d: %d words\n", lineNumber, wordCount)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Найти и заменить слово в файле, с сохранением регистра
// Описание: Напишите программу, которая находит и заменяет все вхождения слова в файле,
// сохраняя регистр (с учетом заглавных и строчных букв).
// Методы или функции:
// func replaceWordWithCase(fileName, oldWord, newWord string) error

func replaceWordWithCase(fileName, oldWord, newWord string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	updatedContent := replaceWithCaseSensitive(string(content), oldWord, newWord)

	return os.WriteFile(fileName, []byte(updatedContent), 0644)
}

func replaceWithCaseSensitive(content, oldWord, newWord string) string {
	words := strings.Fields(content)
	for i, word := range words {
		if strings.EqualFold(word, oldWord) {
			words[i] = matchCase(word, newWord)
		}
	}
	return strings.Join(words, " ")
}

func matchCase(source, target string) string {
	if len(source) == 0 {
		return target
	}

	runes := []rune(target)
	if unicode.IsUpper([]rune(source)[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	} else {
		runes[0] = unicode.ToLower(runes[0])
	}

	return string(runes)
}

// Поиск самого короткого слова в файле
// Описание: Напишите программу, которая находит самое короткое слово в файле и выводит
// его на экран.
// Методы или функции:
// func findShortestWord(fileName string) (string, error)

func findShortestWord(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var shortestWord string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if len(shortestWord) == 0 || len(word) < len(shortestWord) {
			shortestWord = word
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return shortestWord, nil
}

// Чтение и запись файла по частям
// Описание: Напишите программу, которая читает файл по частям (например, по 1024 байта)
// и записывает его содержимое в другой файл.
// Методы или функции:
// func copyFileInChunks(srcFileName, dstFileName string) error

func copyFileInChunks(srcFileName, dstFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := srcFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		_, err = dstFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

// Подсчет количества символов, слов и строк в файле
// Описание: Напишите программу, которая подсчитывает количество символов, слов
// и строк в файле и выводит результат на экран.
// Методы или функции:
// func countFileStats(fileName string) (int, int, int, error)

func countFileStats(fileName string) (int, int, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	charCount, wordCount, lineCount := 0, 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		charCount += len(line)
		wordCount += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return charCount, wordCount, lineCount, nil
}

// Поиск и удаление строки в файле
// Описание: Напишите программу, которая находит и удаляет строку с определенным
// содержимым из файла.
// Методы или функции:
// func deleteLineFromFile(fileName, lineToDelete string) error

func deleteLineFromFile(fileName, lineToDelete string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	tempFileName := "temp_" + fileName
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		line := scanner.Text()
		if line != lineToDelete {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	err = os.Rename(tempFileName, fileName)
	if err != nil {
		return err
	}

	return nil
}

// Сортировка строк в файле в алфавитном порядке
// Описание: Напишите программу, которая сортирует строки в файле в алфавитном порядке
// и записывает результат в новый файл.
// Методы или функции:
// func sortFileLines(fileName, outputFile string) error

func sortFileLines(fileName, outputFile string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Strings(lines)

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}
