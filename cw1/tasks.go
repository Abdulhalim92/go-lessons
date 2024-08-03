package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1. Вариант 1
// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий только
// уникальные элементы из исходного среза.
func unique(slice []int) []int {
	result := []int{}
	seen := map[int]bool{}
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез строк и возвращает срез, в котором строки упорядочены
// по длине. Если длины строк совпадают, то строки должны упорядочиваться в лексикографическом порядке.
func sortByLengthAndLexicographically(slice []string) []string {
	sort.Slice(slice, func(i, j int) bool {
		if len(slice[i]) == len(slice[j]) {
			return slice[i] < slice[j]
		}
		return len(slice[i]) < len(slice[j])
	})
	return slice
}

// • Напишите функцию, которая принимает срез чисел и возвращает срез, где каждое число умножено на 2,
// но при этом сумма всех чисел в исходном срезе должна оставаться постоянной.
func doubleValues(slice []int) []int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	doubled := []int{}
	for _, v := range slice {
		doubled = append(doubled, v*2)
	}
	return doubled
}

// 2. Вариант 2
// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий только
// четные числа из исходного среза, отсортированные по убыванию.
func evenNumbersSorted(slice []int) []int {
	evens := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(evens)))
	return evens
}

// • Напишите функцию, которая принимает срез строк и возвращает строку, состоящую из всех строк
// исходного среза, объединенных в одну, с разделителем "-".
func concatenateWithSeparator(slice []string) string {
	return strings.Join(slice, "-")
}

// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий только
// те строки, которые имеют больше трех символов и начинаются с буквы A.
func filterStrings(slice []string) []string {
	filtered := []string{}
	for _, str := range slice {
		if len(str) > 3 && str[0] == 'A' {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

// 3. Вариант 3
// • Реализуйте функцию, которая принимает срез структур Person и возвращает нового человека,
// у которого имя самое длинное из всех имен в исходном срезе.

type Person struct {
	Name string
	Age  int
}

func longestName(people []Person) Person {
	var longest Person
	for _, p := range people {
		if len(p.Name) > len(longest.Name) {
			longest = p
		}
	}
	return longest
}

// • Напишите функцию, которая принимает срез чисел и возвращает новый срез, содержащий те же числа,
// но только те, которые являются степенями двойки.
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func filterPowersOfTwo(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if isPowerOfTwo(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Rectangle и возвращает структуру Rectangle,
// у которой площадь максимальна.

type Rectangle struct {
	Width  float64
	Height float64
}

func maxAreaRectangle(rectangles []Rectangle) Rectangle {
	var maxRect Rectangle
	maxArea := 0.0
	for _, r := range rectangles {
		area := r.Width * r.Height
		if area > maxArea {
			maxArea = area
			maxRect = r
		}
	}
	return maxRect
}

// 4. Вариант 4
// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий только
// те числа, которые являются простыми.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func filterPrimes(slice []int) []int {
	primes := []int{}
	for _, v := range slice {
		if isPrime(v) {
			primes = append(primes, v)
		}
	}
	return primes
}

// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, где каждая строка
// заменена на её реверс.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStrings(slice []string) []string {
	reversed := []string{}
	for _, str := range slice {
		reversed = append(reversed, reverseString(str))
	}
	return reversed
}

// • Напишите функцию, которая принимает срез строк и возвращает новый срез, содержащий строки,
// длина которых равна числу гласных в этой строке.
func countVowels(s string) int {
	vowels := "aeiou"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func filterByVowelCount(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if len(str) == countVowels(str) {
			result = append(result, str)
		}
	}
	return result
}

// 5. Вариант 5
// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий только
// те строки, которые являются палиндромами.
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == reverseString(s)
}

func filterPalindromes(slice []string) []string {
	palindromes := []string{}
	for _, str := range slice {
		if isPalindrome(str) {
			palindromes = append(palindromes, str)
		}
	}
	return palindromes
}

// • Напишите функцию, которая принимает срез чисел и возвращает срез, содержащий все числа,
// которые являются квадратами целых чисел.
func isSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}

func filterSquares(slice []int) []int {
	squares := []int{}
	for _, v := range slice {
		if isSquare(v) {
			squares = append(squares, v)
		}
	}
	return squares
}

// • Реализуйте функцию, которая принимает срез структур Employee и возвращает структуру
// Employee с максимальной зарплатой.

type Employee struct {
	Name   string
	Salary float64
}

func highestPaidEmployee(employees []Employee) Employee {
	var maxEmp Employee
	maxSalary := 0.0
	for _, emp := range employees {
		if emp.Salary > maxSalary {
			maxSalary = emp.Salary
			maxEmp = emp
		}
	}
	return maxEmp
}

// 6. Вариант 6
// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез,
// содержащий числа, которые встречаются больше одного раза.
func duplicates(slice []int) []int {
	count := map[int]int{}
	for _, v := range slice {
		count[v]++
	}
	duplicates := []int{}
	for k, v := range count {
		if v > 1 {
			duplicates = append(duplicates, k)
		}
	}
	return duplicates
}

// • Напишите функцию, которая принимает срез строк и возвращает срез, содержащий строки,
// которые являются анаграммами друг друга.
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func findAnagrams(slice []string) [][]string {
	anagrams := map[string][]string{}
	for _, str := range slice {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	result := [][]string{}
	for _, group := range anagrams {
		if len(group) > 1 {
			result = append(result, group)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Book и возвращает структуру Book,
// у которой год публикации самый новый.

type Book struct {
	Title string
	Year  int
}

func latestBook(books []Book) Book {
	var latest Book
	latestYear := 0
	for _, b := range books {
		if b.Year > latestYear {
			latestYear = b.Year
			latest = b
		}
	}
	return latest
}

// 7. Вариант 7
// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, в котором
// каждое число возведено в квадрат.
func squareNumbers(slice []int) []int {
	squared := []int{}
	for _, v := range slice {
		squared = append(squared, v*v)
	}
	return squared
}

// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий
// строки в нижнем регистре.
func toLowerCase(slice []string) []string {
	lowerCase := []string{}
	for _, str := range slice {
		lowerCase = append(lowerCase, strings.ToLower(str))
	}
	return lowerCase
}

// • Реализуйте функцию, которая принимает срез структур Car и возвращает структуру Car,
// у которой цена минимальна.

type Car struct {
	Model string
	Price float64
}

func cheapestCar(cars []Car) Car {
	var minCar Car
	minPrice := float64(math.MaxFloat64)
	for _, c := range cars {
		if c.Price < minPrice {
			minPrice = c.Price
			minCar = c
		}
	}
	return minCar
}

// 8. Вариант 8
// • Напишите функцию, которая принимает срез строк и возвращает новый срез, где каждая строка
// заменена на строку с длиной, равной количеству букв в этой строке.
func lengthStrings(slice []string) []string {
	lengths := []string{}
	for _, str := range slice {
		lengths = append(lengths, strings.Repeat("a", len(str)))
	}
	return lengths
}

// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий
// числа, которые можно разделить на 3, но не на 5.
func filterDivisibleBy3Not5(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%3 == 0 && v%5 != 0 {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, где каждая строка
// заменена на её обратный порядок символов.
func reverseString2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStrings2(slice []string) []string {
	reversed := []string{}
	for _, str := range slice {
		reversed = append(reversed, reverseString2(str))
	}
	return reversed
}

// 9. Вариант 9
// • Реализуйте функцию, которая принимает срез структур Student и возвращает структуру Student,
// у которой средний балл наивысший.

type Student struct {
	Name  string
	Grade float64
}

func highestGradeStudent(students []Student) Student {
	var topStudent Student
	maxGrade := 0.0
	for _, s := range students {
		if s.Grade > maxGrade {
			maxGrade = s.Grade
			topStudent = s
		}
	}
	return topStudent
}

// • Напишите функцию, которая принимает срез строк и возвращает строку, содержащую все уникальные
// символы, встречающиеся в срезе строк.
func uniqueCharacters(slice []string) string {
	chars := map[rune]bool{}
	for _, str := range slice {
		for _, char := range str {
			chars[char] = true
		}
	}
	result := []rune{}
	for char := range chars {
		result = append(result, char)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return string(result)
}

// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез, где каждое
// число возведено в куб.
func cubeNumbers(slice []int) []int {
	cubed := []int{}
	for _, v := range slice {
		cubed = append(cubed, v*v*v)
	}
	return cubed
}

// 10. Вариант 10
// • Реализуйте функцию, которая принимает срез строк и возвращает строку, состоящую из всех строк
// исходного среза, объединенных через пробел.
func joinWithSpace(slice []string) string {
	return strings.Join(slice, " ")
}

// • Напишите функцию, которая принимает срез целых чисел и возвращает срез, содержащий только те
// числа, которые делятся на 7 и не делятся на 3.
func filterDivisibleBy7Not3(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%7 == 0 && v%3 != 0 {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Product и возвращает структуру Product,
// у которой цена наивысшая.

type Product struct {
	Name  string
	Price float64
}

func highestPricedProduct(products []Product) Product {
	var maxProduct Product
	maxPrice := 0.0
	for _, p := range products {
		if p.Price > maxPrice {
			maxPrice = p.Price
			maxProduct = p
		}
	}
	return maxProduct
}

// 11. Вариант 11
// • Напишите функцию, которая принимает срез строк и возвращает срез, где каждая строка заменена
// на её длину в виде строки.
func stringLengths(slice []string) []string {
	lengths := []string{}
	for _, str := range slice {
		lengths = append(lengths, strconv.Itoa(len(str)))
	}
	return lengths
}

// • Реализуйте функцию, которая принимает срез целых чисел и возвращает срез, содержащий числа,
// которые являются чётными квадратами.
func isEvenSquare(n int) bool {
	if n%2 != 0 {
		return false
	}
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}

func filterEvenSquares(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if isEvenSquare(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Transaction и возвращает сумму всех
// транзакций.

type Transaction struct {
	Amount float64
}

func totalAmount(transactions []Transaction) float64 {
	sum := 0.0
	for _, t := range transactions {
		sum += t.Amount
	}
	return sum
}

// 12. Вариант 12
// • Реализуйте функцию, которая принимает срез строк и возвращает срез, содержащий только те
// строки, которые содержат букву e.
func containsE(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if strings.Contains(str, "e") {
			result = append(result, str)
		}
	}
	return result
}

// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий
// только те числа, которые являются неотрицательными.
func filterNonNegative(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v >= 0 {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Movie и возвращает структуру Movie,
// у которой рейтинг наивысший.

type Movie struct {
	Title  string
	Rating float64
}

func highestRatedMovie(movies []Movie) Movie {
	var topMovie Movie
	maxRating := 0.0
	for _, m := range movies {
		if m.Rating > maxRating {
			maxRating = m.Rating
			topMovie = m
		}
	}
	return topMovie
}

// 13. Вариант 13
// • Реализуйте функцию, которая принимает срез строк и возвращает строку, состоящую из всех
// строк в обратном порядке.
func reverseOrder(slice []string) string {
	reversed := []string{}
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return strings.Join(reversed, " ")
}

// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий
// числа, которые являются нечетными кубами.
func isOddCube(n int) bool {
	if n%2 == 0 {
		return false
	}
	cubeRoot := math.Cbrt(float64(n))
	return cubeRoot == float64(int(cubeRoot))
}

func filterOddCubes(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if isOddCube(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Order и возвращает структуру Order
// с наименьшим количеством единиц продукта.

type Order struct {
	Product  string
	Quantity int
}

func smallestOrder(orders []Order) Order {
	var minOrder Order
	minQuantity := int(^uint(0) >> 1) // Максимально возможное значение для int
	for _, o := range orders {
		if o.Quantity < minQuantity {
			minQuantity = o.Quantity
			minOrder = o
		}
	}
	return minOrder
}

// 14. Вариант 14
// • Реализуйте функцию, которая принимает срез структур Player и возвращает структуру Player с
// наибольшим количеством очков.

type Player struct {
	Name  string
	Score int
}

func highestScoringPlayer(players []Player) Player {
	var topPlayer Player
	maxScore := 0
	for _, p := range players {
		if p.Score > maxScore {
			maxScore = p.Score
			topPlayer = p
		}
	}
	return topPlayer
}

// • Напишите функцию, которая принимает срез строк и возвращает срез, содержащий только те строки,
// которые не содержат пробелы
func noSpaces(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if !strings.Contains(str, " ") {
			result = append(result, str)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Product и возвращает сумму всех цен.

//type Product struct {
//	Name  string
//	Price float64
//}

func totalPrice(products []Product) float64 {
	sum := 0.0
	for _, p := range products {
		sum += p.Price
	}
	return sum
}

// 15. Вариант 15
// • Напишите функцию, которая принимает срез строк и возвращает срез, содержащий строки,
// отсортированные по длине.

func sortByLength(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})
	return sorted
}

// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий
// только те числа, которые не являются простыми.

func filterNonPrimes(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if !isPrime(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Book и возвращает срез, отсортированный по
// году публикации в порядке убывания.

//type Book struct {
//	Title string
//	Year  int
//}

func sortBooksByYear(books []Book) []Book {
	sorted := make([]Book, len(books))
	copy(sorted, books)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Year > sorted[j].Year
	})
	return sorted
}

// 16. Вариант 16
// • Реализуйте функцию, которая принимает срез структур Event и возвращает структуру Event, у
// которой продолжительность самая длинная.

type Event struct {
	Name     string
	Duration int // Продолжительность в минутах
}

func longestEvent(events []Event) Event {
	var longest Event
	maxDuration := 0
	for _, e := range events {
		if e.Duration > maxDuration {
			maxDuration = e.Duration
			longest = e
		}
	}
	return longest
}

// • Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий
// только те числа, которые являются квадратами чисел.

func isSquare2(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := math.Sqrt(float64(n))
	return sqrt == float64(int(sqrt))
}

func filterSquares2(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if isSquare2(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, содержащий только
// те строки, которые являются палиндромами.
func isPalindrome3(s string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	reversed := reverseString3(cleaned)
	return cleaned == reversed
}

func reverseString3(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func filterPalindromes3(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if isPalindrome3(str) {
			result = append(result, str)
		}
	}
	return result
}

// 17. Вариант 17
// • Напишите функцию, которая принимает срез структур Animal и возвращает структуру Animal, у которой
// средний возраст наибольший.

type Animal struct {
	Name string
	Age  int
}

func oldestAnimal(animals []Animal) Animal {
	var oldest Animal
	maxAge := 0
	for _, a := range animals {
		if a.Age > maxAge {
			maxAge = a.Age
			oldest = a
		}
	}
	return oldest
}

// • Реализуйте функцию, которая принимает срез строк и возвращает срез, где каждая строка заменена
// на строку с удаленными все вхождениями символа a.
func removeCharA(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		result = append(result, strings.ReplaceAll(str, "a", ""))
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Transaction и возвращает новый срез,
// содержащий транзакции, сумма которых превышает заданное значение.

//type Transaction struct {
//	Amount float64
//}

func filterTransactions(transactions []Transaction, threshold float64) []Transaction {
	result := []Transaction{}
	for _, t := range transactions {
		if t.Amount > threshold {
			result = append(result, t)
		}
	}
	return result
}

// 18. Вариант 18
// • Реализуйте функцию, которая принимает срез строк и возвращает новый срез, где каждая строка
// является перевернутой строкой из исходного среза.
func reverseStrings4(slice []string) []string {
	reversed := []string{}
	for _, str := range slice {
		reversed = append(reversed, reverseString(str))
	}
	return reversed
}

// • Напишите функцию, которая принимает срез целых чисел и возвращает срез, содержащий только те числа,
// которые являются степенями двойки.
func isPowerOfTwo2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func filterPowersOfTwo2(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if isPowerOfTwo2(v) {
			result = append(result, v)
		}
	}
	return result
}

// • Реализуйте функцию, которая принимает срез структур Person и возвращает структуру Person
// с наименьшим возрастом.

//type Person struct {
//	Name string
//	Age  int
//}

func youngestPerson(people []Person) Person {
	var youngest Person
	minAge := int(^uint(0) >> 1) // Максимально возможное значение для int
	for _, p := range people {
		if p.Age < minAge {
			minAge = p.Age
			youngest = p
		}
	}
	return youngest
}

// 19. Вариант 19
// • Реализуйте функцию, которая принимает срез структур Order и возвращает новый срез,
// содержащий только те заказы, у которых сумма товаров превышает заданное значение.

type Order2 struct {
	Product  string
	Quantity int
	Price    float64
}

func filterOrdersByTotal(orders []Order2, minTotal float64) []Order2 {
	result := []Order2{}
	for _, o := range orders {
		total := float64(o.Quantity) * o.Price
		if total > minTotal {
			result = append(result, o)
		}
	}
	return result
}

// • Напишите функцию, которая принимает срез строк и возвращает строку, в которой каждая
// строка исходного среза повторяется трижды, разделенная пробелами.
func repeatStrings(slice []string) string {
	result := []string{}
	for _, str := range slice {
		repeated := strings.Repeat(str, 3)
		result = append(result, repeated)
	}
	return strings.Join(result, " ")
}

// • Реализуйте функцию, которая принимает срез целых чисел и возвращает новый срез,
// содержащий только числа, которые являются простыми и меньше заданного значения.
func filterPrimesLessThan(slice []int, max int) []int {
	result := []int{}
	for _, v := range slice {
		if isPrime(v) && v < max {
			result = append(result, v)
		}
	}
	return result
}
