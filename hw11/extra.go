package main

import (
	"strings"
	"unicode"
)

// Обработка строк с использованием методов и callback
// Создайте структуру StringProcessor, которая содержит метод для обработки строки с использованием
// функции callback. Функция callback должна применяться к каждой строке. Реализуйте несколько
// функций callback для преобразования строки, таких как конвертация в верхний регистр и удаление
// пробелов.

// StringProcessor содержит метод для обработки строки.
type StringProcessor struct {
	Content string
}

// Process применяет функцию callback к строке.
func (sp *StringProcessor) Process(callback func(string) string) string {
	return callback(sp.Content)
}

// Форматирование строки с использованием callback
// Создайте структуру Formatter, которая имеет метод для форматирования строки. Используйте
// callback для определения способа форматирования строки. Реализуйте различные способы
// форматирования, такие как добавление префикса и суффикса.

// Formatter содержит строку и метод для её форматирования.
type Formatter struct {
	Content string
}

// NewFormatter возвращает новый Formatter с заданным форматирующим closure.
func NewFormatter(s string, formatter func(string) string) *Formatter {
	return &Formatter{Content: formatter(s)}
}

// Обработка текста с использованием методов и callback
// Создайте структуру TextProcessor, которая содержит метод для обработки текста с использованием
// функции callback. Функция callback должна возвращать количество слов в строке. Реализуйте
// функцию callback для подсчета количества слов и символов.

// TextProcessor содержит метод для обработки текста.
type TextProcessor struct {
	Content string
}

// Process применяет функцию callback к строке и возвращает результат.
func (tp *TextProcessor) Process(callback func(string) int) int {
	return callback(tp.Content)
}

func countWords2(s string) int {
	return len(strings.Fields(s))
}

func countCharacters2(s string) int {
	return len(s)
}

// Изменение строки с использованием методов и callback
// Создайте структуру StringManipulator, которая содержит метод для изменения строки с
// использованием callback. Реализуйте несколько закрытых функций для преобразования строки,
// таких как инверсия строки и замена пробелов на тире.

// StringManipulator содержит строку и метод для её изменения.
type StringManipulator struct {
	Content string
}

// Manipulate применяет closure к строке и возвращает результат.
func (sm *StringManipulator) Manipulate(manipulator func(string) string) string {
	return manipulator(sm.Content)
}

//func reverseString(s string) string {
//	runes := []rune(s)
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//	return string(runes)
//}

func replaceSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}

// Обработка строки с использованием методов и callback для фильтрации
// Создайте структуру StringFilter, которая имеет метод для фильтрации строки с
// использованием функции callback. Функция callback должна проверять, удовлетворяет ли
// строка определенному условию. Реализуйте функции callback для проверки наличия подстроки
// и длины строки.

// StringFilter содержит метод для фильтрации строки.
type StringFilter struct {
	Content string
}

// Filter применяет функцию callback к строке и возвращает результат.
func (sf *StringFilter) Filter(callback func(string) bool) bool {
	return callback(sf.Content)
}

func containsWord(s string, word string) bool {
	return strings.Contains(s, word)
}

func isLongerThan(s string, length int) bool {
	return len(s) > length
}

// Изменение строки с использованием методов и callback
// Создайте структуру TextEditor, которая содержит метод для изменения текста с использованием
// функции callback. Реализуйте несколько функций callback для изменения текста, таких как
// добавление текста в начало и конец строки.

// TextEditor содержит текст и метод для его изменения.
type TextEditor struct {
	Content string
}

// Edit применяет функцию callback к тексту.
func (te *TextEditor) Edit(callback func(string) string) {
	te.Content = callback(te.Content)
}

func addPrefix(s string) string {
	return "Prefix: " + s
}

func addSuffix(s string) string {
	return s + " :Suffix"
}

// Обработка строки с использованием методов и closure для статистики
// Создайте структуру Statistics, которая содержит метод для подсчета статистики строки.
// Используйте closure для подсчета различных метрик, таких как количество пробелов и
// количество уникальных символов.

// Statistics содержит текст и метод для подсчета статистики.
type Statistics struct {
	Content string
}

// Compute применяет closure для вычисления статистики.
func (st *Statistics) Compute(statFunc func(string) int) int {
	return statFunc(st.Content)
}

func countSpaces(s string) int {
	return strings.Count(s, " ")
}

func countUniqueCharacters(s string) int {
	unique := make(map[rune]bool)
	for _, r := range s {
		unique[r] = true
	}
	return len(unique)
}

// Замена строки с использованием методов и callback
// Создайте структуру Replacer, которая содержит метод для замены строки с использованием
// функции callback. Реализуйте функции callback для замены всех вхождений подстроки и
// удаления всех цифр.

// Replacer содержит строку и метод для её замены.
type Replacer struct {
	Content string
}

// Replace применяет функцию callback к строке.
func (r *Replacer) Replace(callback func(string) string) {
	r.Content = callback(r.Content)
}

//func replaceSubstring(s string, old, new string) string {
//	return strings.ReplaceAll(s, old, new)
//}

func removeDigits(s string) string {
	result := ""
	for _, r := range s {
		if !unicode.IsDigit(r) {
			result += string(r)
		}
	}
	return result
}

// Форматирование строки с использованием методов и closure
// Создайте структуру TextFormatter, которая содержит метод для форматирования строки с
// использованием closure. Реализуйте разные функции форматирования, такие как подчеркивание
// и замена пробелов на подчеркивания.

// TextFormatter содержит строку и метод для её форматирования.
type TextFormatter struct {
	Content string
}

// Format применяет closure для форматирования строки.
func (tf *TextFormatter) Format(formatter func(string) string) {
	tf.Content = formatter(tf.Content)
}

func underline(s string) string {
	return strings.Join(strings.Split(s, ""), "_")
}

func replaceSpacesWithUnderscore(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// Анализ строки с использованием методов и callback
// Создайте структуру StringAnalyzer, которая содержит метод для анализа строки с использованием
// функции callback. Реализуйте функции callback для подсчета количества гласных и согласных букв.

// StringAnalyzer содержит текст и метод для его анализа.
type StringAnalyzer struct {
	Content string
}

// Analyze применяет функцию callback для анализа строки.
func (sa *StringAnalyzer) Analyze(callback func(string) int) int {
	return callback(sa.Content)
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			count++
		}
	}
	return count
}

func countConsonants(s string) int {
	consonants := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	count := 0
	for _, r := range s {
		if strings.ContainsRune(consonants, r) {
			count++
		}
	}
	return count
}
