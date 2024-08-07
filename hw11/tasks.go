package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// // Конкатенация строк
// Напишите функцию, которая принимает две строки и возвращает их конкатенацию.
func concatenateStrings(s1, s2 string) string {
	return s1 + s2
}

// Длина строки
// Напишите функцию, которая принимает строку и возвращает её длину.
func stringLength(s string) int {
	return len(s)
}

// Проверка подстроки
// Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.
func containsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Поиск подстроки
// Напишите функцию, которая находит индекс первого вхождения подстроки в строке.
// Верните -1, если подстрока не найдена.
func indexOfSubstring(s, substr string) int {
	return strings.Index(s, substr)
}

// Замена подстроки
// Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку.
func replaceSubstring(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Удаление пробелов
// Напишите функцию, которая удаляет пробелы в начале и в конце строки.
func trimSpaces(s string) string {
	return strings.TrimSpace(s)
}

// Преобразование регистра
// Напишите функцию, которая преобразует строку к верхнему и нижнему регистру.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

// Повторение строки
// Напишите функцию, которая повторяет строку заданное количество раз.
func repeatString(s string, count int) string {
	return strings.Repeat(s, count)
}

// Разбиение строки
// Напишите функцию, которая разбивает строку по заданному разделителю и возвращает срез строк.
func splitString(s, sep string) []string {
	return strings.Split(s, sep)
}

// Сравнение строк
// Напишите функцию, которая сравнивает две строки без учета регистра.
func equalFold(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// Поиск и замена первой подстроки
// Напишите функцию, которая заменяет первое вхождение подстроки в строке на другую подстроку.
func replaceFirst(s, old, new string) string {
	return strings.Replace(s, old, new, 1)
}

// Инверсия строки
// Напишите функцию, которая переворачивает строку.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Подсчет символов
// Напишите функцию, которая подсчитывает количество каждого символа в строке и возвращает
// карту с результатами.
func countCharacters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

// Удаление символов
// Напишите функцию, которая удаляет все вхождения заданного символа из строки.
func removeCharacter(s string, char rune) string {
	return strings.ReplaceAll(s, string(char), "")
}

// Подсчет слов
// Напишите функцию, которая подсчитывает количество слов в строке. Словом считается
// последовательность символов, разделенная пробелами.
func countWords(s string) int {
	return len(strings.Fields(s))
}

// Проверка префикса и суффикса
// Напишите функции, которые проверяют, начинается ли строка с заданного префикса и
// заканчивается ли строка заданным суффиксом.
func hasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func hasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Удаление дубликатов символов
// Напишите функцию, которая удаляет дубликаты символов в строке.
//func removeDuplicates(s string) string {
//	seen := make(map[rune]bool)
//	result := make([]rune, 0, len(s))
//	for _, r := range s {
//		if !seen[r] {
//			seen[r] = true
//			result = append(result, r)
//		}
//	}
//	return string(result)
//}

func removeDuplicates(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		if !contains(result, r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func contains(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}

// Форматирование строки
// Напишите функцию, которая форматирует строку, заменяя специальные символы HTML на их
// эквиваленты.
func escapeHTML(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(s)
}

// Псевдонимы строк
// Напишите функцию, которая создает псевдоним для строки, заменяя пробелы и специальные
// символы подчеркиваниями.
func createAlias(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	alias := re.ReplaceAllString(s, "_")
	return alias
}

// Разбор строки
// Напишите функцию, которая разбирает строку, содержащую числа, и возвращает их сумму.
func sumNumbers(s string) int {
	parts := strings.Fields(s)
	sum := 0
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			sum += num
		}
	}
	return sum
}

// Обратный порядок слов
// Напишите функцию, которая принимает строку и возвращает строку с обратным порядком слов.
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// Подсчет уникальных слов
// Напишите функцию, которая подсчитывает количество уникальных слов в строке.
//func countUniqueWords(s string) int {
//	words := strings.Fields(s)
//	uniqueWords := make(map[string]bool)
//	for _, word := range words {
//		uniqueWords[word] = true
//	}
//	return len(uniqueWords)
//}

func countUniqueWords(s string) int {
	words := strings.Fields(s)
	uniqueWords := make([]string, 0, len(words))

	for _, word := range words {
		if !contains2(uniqueWords, word) {
			uniqueWords = append(uniqueWords, word)
		}
	}

	return len(uniqueWords)
}

func contains2(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// Поиск палиндромов
// Напишите функцию, которая проверяет, является ли строка палиндромом.
func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// Перемешивание слов
// Напишите функцию, которая случайным образом перемешивает слова в строке.
func shuffleWords(s string) string {
	words := strings.Fields(s)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return strings.Join(words, " ")
}

// Сортировка слов по длине
// Напишите функцию, которая сортирует слова в строке по их длине.
func sortByLength(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	return strings.Join(words, " ")
}

// Генерация хеш-значения
// Напишите функцию, которая вычисляет хеш-значение строки с использованием алгоритма SHA-256.
func hashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// Генерация всех подстрок
// Напишите функцию, которая генерирует все подстроки заданной строки.
func allSubstrings(s string) []string {
	substrings := []string{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substrings = append(substrings, s[i:j])
		}
	}
	return substrings
}

// Поиск всех анаграмм
// Напишите функцию, которая ищет все анаграммы строки в другой строке.
func findAllAnagrams(s, t string) []string {
	sortedT := sortString(t)
	anagrams := []string{}
	for i := 0; i <= len(s)-len(t); i++ {
		if sortedT == sortString(s[i:i+len(t)]) {
			anagrams = append(anagrams, s[i:i+len(t)])
		}
	}
	return anagrams
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Подсчет слов и символов
// Напишите функцию, которая подсчитывает количество слов и символов в строке.
func countWordsAndCharacters(s string) (int, int) {
	words := strings.Fields(s)
	characters := len(s)
	return len(words), characters
}
