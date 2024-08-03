package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {

}

func workWithStrings() {
	// Иммутабельность (неизменяемость) строк
	s := "hello"
	s = "world"
	fmt.Println(s) // "world"

	// Конкатенация строк
	s1 := "Hello"
	s2 := "World"
	s3 := s1 + " " + s2
	fmt.Println(s3) // "Hello World"

	parts := []string{"Hello", "World"}
	s4 := strings.Join(parts, " ")
	fmt.Println(s4) // "Hello World"

	// Форматирование строк
	s = fmt.Sprintf("Hello, %s!", "World")
	fmt.Println(s) // "Hello, World!"

	// Сравнение строк
	s1 = "apple"
	s2 = "banana"
	fmt.Println(s1 == s2) // false
	fmt.Println(s1 < s2)  // true

	// Разбиение строки
	s = "a,b,c"
	parts = strings.Split(s, ",")
	fmt.Println(parts) // ["a", "b", "c"]

	// Разбиение на максимум 2 части
	parts = strings.SplitN(s, ",", 2)
	fmt.Println(parts) // ["a", "b,c"]

	// Поиск подстроки
	s = "Hello, World!"
	fmt.Println(strings.Contains(s, "World")) // true
	fmt.Println(strings.Index(s, "World"))    // 7
	fmt.Println(strings.LastIndex(s, "o"))    // 8

	// Удаление пробелов и изменение регистра
	s = "  Hello, World!  "
	fmt.Println(strings.TrimSpace(s)) // "Hello, World!"
	fmt.Println(strings.ToUpper(s))   // "  HELLO, WORLD!  "
	fmt.Println(strings.ToLower(s))   // "  hello, world!  "

	// Проверка на равенство с игнорированием регистра
	s1 = "Hello"
	s2 = "hello"
	fmt.Println(strings.EqualFold(s1, s2)) // true

	// Извлечение подстрок
	s = "Hello, World!"
	sub := s[7:12]
	fmt.Println(sub) // "World"

	// Изменение строки на основе шаблонов
	s = "Hello, World!"
	newS := strings.Replace(s, "World", "Go", -1)
	fmt.Println(newS) // "Hello, Go!"
}

func useStrings() {
	// Подсчет вхождений подстроки в строку
	s := "banana"
	count := strings.Count(s, "a")
	fmt.Println(count) // 3

	// Проверка наличия подстроки в строке
	s = "Hello, World!"
	contains := strings.Contains(s, "World")
	fmt.Println(contains) // true

	// Разделение строки на подстроки
	s = "a,b,c"
	split := strings.Split(s, ",")
	fmt.Println(split) // ["a", "b", "c"]

	// Объединение строк с разделителем
	parts := []string{"Go", "is", "great"}
	joined := strings.Join(parts, " ")
	fmt.Println(joined) // "Go is great"

	// Замена подстроки
	s = "Hello, Gophers!"
	replaced := strings.Replace(s, "Gophers", "World", -1)
	fmt.Println(replaced) // "Hello, World!"
}

func useStrConv() {
	// Преобразование строки в целое число
	s := "42"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i) // 42
	}

	// Преобразование целого числа в строку
	i = 42
	s = strconv.Itoa(i)
	fmt.Println(s) // "42"

	// Преобразование строки в число с плавающей запятой
	s = "3.14"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f) // 3.14
	}

	// Преобразование числа с плавающей запятой в строку
	f = 3.14
	s = strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println(s) // "3.14"
}

func useUnicode() {
	// Проверка, является ли символ буквой или цифрой
	r := 'A'
	fmt.Println(unicode.IsLetter(r)) // true
	r = '1'
	fmt.Println(unicode.IsDigit(r)) // true

	// Преобразование строки в верхний регистр
	s := "Hello, World!"
	upper := toUpper(s)
	fmt.Println(upper) // "HELLO, WORLD!"

	// Проверка принадлежности символа к нескольким категориям
	r = 'A'
	fmt.Println(unicode.IsLetter(r)) // true
	fmt.Println(unicode.IsDigit(r))  // false
}

func toUpper(s string) string {
	result := []rune{}
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		result = append(result, unicode.ToUpper(r))
		s = s[size:]
	}
	return string(result)
}

func useUtf8() {
	// Декодирование рун из строки
	s := "Hello, 世界"
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("%c ", r) // H e l l o ,   世 界
		s = s[size:]
	}

	// Подсчет количества рун в строке
	s = "Hello, 世界"
	count := utf8.RuneCountInString(s)
	fmt.Printf("Number of runes: %d\n", count) // Number of runes: 9

	// Проверка допустимости UTF-8 строки
	s = "Hello, 世界"
	valid := utf8.ValidString(s)
	fmt.Printf("String valid: %t\n", valid) // String valid: true

	// Кодирование и декодирование рун
	r := '世'
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Printf("Encoded: % x\n", buf) // Encoded: e4 b8 96
	decodedRune, _ := utf8.DecodeRune(buf)
	fmt.Printf("Decoded rune: %c\n", decodedRune) // Decoded rune: 世
}
