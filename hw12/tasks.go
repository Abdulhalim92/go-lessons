package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Создание и вывод map
// Описание: Создайте map для хранения строк и их длин, добавьте несколько
// элементов и выведите содержимое.
func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// Проверка наличия ключа
// Описание: Создайте map с несколькими элементами и напишите функцию,
// которая проверяет наличие заданного ключа.
func keyExists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// Обновление значения по ключу
// Описание: Напишите функцию для обновления значения в map по заданному ключу.
func updateValue(m map[string]int, key string, value int) {
	m[key] = value
}

// Удаление элемента из map
// Описание: Создайте функцию, которая удаляет элемент из map по заданному ключу.
func removeKey(m map[string]int, key string) {
	delete(m, key)
}

// Подсчет частоты слов
// Описание: Напишите функцию, которая подсчитывает частоту слов в строке и
// возвращает map с результатами.
func countWords(s string) map[string]int {
	words := strings.Fields(s)
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

// Сортировка ключей в map
// Описание: Напишите функцию, которая сортирует ключи в map и возвращает
// отсортированные ключи.
func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Проверка пустого map
// Описание: Напишите функцию, которая проверяет, пустой ли map.
func isEmpty(m map[string]int) bool {
	return len(m) == 0
}

// Инвертирование ключей и значений
// Описание: Напишите функцию, которая инвертирует ключи и значения в map.
func invertMap(m map[string]int) map[int]string {
	inverted := make(map[int]string)
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}

// Проверка дубликатов в map
// Описание: Напишите функцию, которая проверяет, есть ли дубликаты значений в map.
func hasDuplicates(m map[string]int) bool {
	seen := make(map[int]bool)
	for _, value := range m {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}

// Получение всех значений
// Описание: Напишите функцию, которая возвращает все значения из map в срезе.
func values(m map[string]int) []int {
	vals := make([]int, 0, len(m))
	for _, value := range m {
		vals = append(vals, value)
	}
	return vals
}

// Подсчет уникальных значений в срезе строк
// Описание: Напишите функцию, которая подсчитывает уникальные значения в
// срезе строк с использованием map.
func uniqueValues(slice []string) map[string]bool {
	unique := make(map[string]bool)
	for _, value := range slice {
		unique[value] = true
	}
	return unique
}

// Поиск значений по маске
// Описание: Напишите функцию, которая возвращает все ключи, чьи значения
// удовлетворяют заданному условию.
func filterMap(m map[string]int, condition func(int) bool) map[string]int {
	filtered := make(map[string]int)
	for key, value := range m {
		if condition(value) {
			filtered[key] = value
		}
	}
	return filtered
}

// Объединение двух map
// Описание: Напишите функцию для объединения двух map. Если ключи совпадают,
// значения из второго map должны заменять значения из первого.
func mergeMaps(m1, m2 map[string]int) map[string]int {
	merged := make(map[string]int)
	for key, value := range m1 {
		merged[key] = value
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

// Поиск по значению и получение ключа
// Описание: Напишите функцию, которая возвращает ключ по значению в map.
// Если значение не найдено, верните пустую строку.
func findKeyByValue(m map[string]int, value int) string {
	for key, v := range m {
		if v == value {
			return key
		}
	}
	return ""
}

// Пересечение двух map
// Описание: Напишите функцию для нахождения пересечения двух map на основе
// ключей и значений.
func intersectMaps(m1, m2 map[string]int) map[string]int {
	intersected := make(map[string]int)
	for key, value1 := range m1 {
		if value2, exists := m2[key]; exists && value1 == value2 {
			intersected[key] = value1
		}
	}
	return intersected
}

// Создание среза ключей из map
// Описание: Напишите функцию, которая возвращает срез всех ключей из map.
func mapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Упаковка значений map в строки
// Описание: Напишите функцию, которая упаковывает все значения map в одну
// строку, разделенную запятой.
func joinValues(m map[string]int) string {
	values := make([]string, 0, len(m))
	for _, value := range m {
		values = append(values, fmt.Sprintf("%d", value))
	}
	return strings.Join(values, ", ")
}

// Копирование map
// Описание: Напишите функцию, которая создает копию map.
func copyMap(m map[string]int) map[string]int {
	copyMaps := make(map[string]int)
	for key, value := range m {
		copyMaps[key] = value
	}
	return copyMaps
}

// Преобразование map в срез пар ключ-значение
// Описание: Напишите функцию, которая преобразует map в срез пар
// ключ-значение.
func mapToSlice(m map[string]int) []struct {
	Key   string
	Value int
} {
	slice := make([]struct {
		Key   string
		Value int
	}, 0, len(m))
	for key, value := range m {
		slice = append(slice, struct {
			Key   string
			Value int
		}{key, value})
	}
	return slice
}

// Обновление значений map на основе другой функции
// Описание: Напишите функцию, которая обновляет значения в map, применяя
// функцию к каждому значению.
func updateValues(m map[string]int, updater func(int) int) {
	for key, value := range m {
		m[key] = updater(value)
	}
}

// Кэширование с использованием map и sync.RWMutex
// Описание: Реализуйте простой кэш с использованием map и sync.RWMutex.

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Сложный фильтр по нескольким условиям
// Описание: Напишите функцию, которая фильтрует map по нескольким условиям.
func complexFilter(m map[string]int, conditions ...func(int) bool) map[string]int {
	filtered := make(map[string]int)
	for key, value := range m {
		meetsAllConditions := true
		for _, condition := range conditions {
			if !condition(value) {
				meetsAllConditions = false
				break
			}
		}
		if meetsAllConditions {
			filtered[key] = value
		}
	}
	return filtered
}

// Слияние и агрегация значений из нескольких map
// Описание: Реализуйте функцию, которая сливает несколько map в один, суммируя
// значения для одинаковых ключей.
func mergeAndAggregate(maps ...map[string]int) map[string]int {
	merged := make(map[string]int)
	for _, m := range maps {
		for key, value := range m {
			merged[key] += value
		}
	}
	return merged
}

// Итерирование по ключам и значениям с использованием функции обратного вызова
// Описание: Напишите функцию, которая итерирует по map, вызывая функцию обратного
// вызова для каждого ключа и значения.
func iterateMap(m map[string]int, callback func(string, int)) {
	for key, value := range m {
		callback(key, value)
	}
}

// Создание и использование карты с указателями на структуры
// Описание: Создайте map с указателями на структуры в качестве значений.

type Person struct {
	Name string
	Age  int
}

// Сопоставление map и массивов
// Описание: Напишите функцию, которая преобразует map в массив структур и наоборот.

type Entry struct {
	Key   string
	Value int
}

func mapToArray(m map[string]int) []Entry {
	array := make([]Entry, 0, len(m))
	for key, value := range m {
		array = append(array, Entry{Key: key, Value: value})
	}
	return array
}

func arrayToMap(arr []Entry) map[string]int {
	m := make(map[string]int)
	for _, entry := range arr {
		m[entry.Key] = entry.Value
	}
	return m
}

// Реализация метода для структуры с map
// Описание: Реализуйте метод для структуры, который работает с map в этой структуре.

type StudentGrades struct {
	grades map[string]int
}

func (sg *StudentGrades) AddGrade(subject string, grade int) {
	sg.grades[subject] = grade
}

func (sg *StudentGrades) AverageGrade() float64 {
	total := 0
	count := 0
	for _, grade := range sg.grades {
		total += grade
		count++
	}
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count)
}

// Использование map для реализации простого хранилища данных
// Описание: Реализуйте простое хранилище данных с map для управления записями.

type Record struct {
	ID   int
	Name string
}

type Storage struct {
	records map[int]Record
}

func (s *Storage) AddRecord(id int, name string) {
	s.records[id] = Record{ID: id, Name: name}
}

func (s *Storage) GetRecord(id int) (Record, bool) {
	record, exists := s.records[id]
	return record, exists
}

// Реализация кэширования с лимитом на количество элементов
// Описание: Реализуйте кэш с использованием map и ограничением на количество
// элементов.

type LimitedCache struct {
	data     map[string]string
	capacity int
	keys     []string
}

func NewLimitedCache(capacity int) *LimitedCache {
	return &LimitedCache{
		data:     make(map[string]string),
		capacity: capacity,
	}
}

func (c *LimitedCache) Set(key, value string) {
	if len(c.data) >= c.capacity {
		oldestKey := c.keys[0]
		c.keys = c.keys[1:]
		delete(c.data, oldestKey)
	}
	c.data[key] = value
	c.keys = append(c.keys, key)
}

func (c *LimitedCache) Get(key string) (string, bool) {
	value, exists := c.data[key]
	return value, exists
}

// Реализация map для хранения срезов строк
// Описание: Реализуйте map, где ключи – это строки, а значения – срезы строк.
// Добавьте функции для добавления, удаления и получения срезов.

type StringMap struct {
	data map[string][]string
}

func (sm *StringMap) Add(key string, values []string) {
	sm.data[key] = values
}

func (sm *StringMap) Remove(key string) {
	delete(sm.data, key)
}

func (sm *StringMap) Get(key string) []string {
	return sm.data[key]
}
