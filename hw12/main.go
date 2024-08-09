package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = len("apple")
	m["banana"] = len("banana")
	printMap(m)

	m2 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println("apple exists:", keyExists(m2, "apple"))
	fmt.Println("grape exists:", keyExists(m2, "grape"))

	m3 := map[string]int{"apple": 1, "banana": 2}
	updateValue(m3, "banana", 3)
	fmt.Println(m3)

	m4 := map[string]int{"apple": 1, "banana": 2}
	removeKey(m4, "apple")
	fmt.Println(m4)

	s := "hello world hello"
	fmt.Println(countWords(s))

	m5 := map[string]int{"banana": 1, "apple": 2}
	fmt.Println(sortedKeys(m5))

	m6 := map[string]int{}
	m7 := map[string]int{"key": 1}
	fmt.Println("m1 empty:", isEmpty(m6))
	fmt.Println("m2 empty:", isEmpty(m7))

	m8 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(invertMap(m8))

	m9 := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println("Has duplicates:", hasDuplicates(m9))

	m10 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(values(m10))

	slice := []string{"a", "b", "a", "c", "b"}
	fmt.Println(uniqueValues(slice))

	m11 := map[string]int{"apple": 5, "banana": 10, "orange": 15}
	filtered := filterMap(m11, func(value int) bool {
		return value > 10
	})
	fmt.Println(filtered)

	m12 := map[string]int{"a": 1, "b": 2}
	m13 := map[string]int{"b": 3, "c": 4}
	fmt.Println(mergeMaps(m12, m13))

	m14 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println("Key for value 1:", findKeyByValue(m14, 1))
	fmt.Println("Key for value 3:", findKeyByValue(m14, 3))

	m15 := map[string]int{"a": 1, "b": 2, "c": 3}
	m16 := map[string]int{"b": 2, "c": 4}
	fmt.Println(intersectMaps(m15, m16))

	m17 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(mapKeys(m17))

	m18 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(joinValues(m18))

	m19 := map[string]int{"apple": 1, "banana": 2}
	copyMap := copyMap(m19)
	fmt.Println(copyMap)

	m20 := map[string]int{"apple": 1, "banana": 2}
	fmt.Println(mapToSlice(m20))

	m21 := map[string]int{"apple": 1, "banana": 2}
	updateValues(m21, func(value int) int { return value * 2 })
	fmt.Println(m21)

	cache := Cache{data: make(map[string]string)}
	cache.Set("foo", "bar")
	if value, exists := cache.Get("foo"); exists {
		fmt.Println("Value:", value)
	}

	m22 := map[string]int{"apple": 5, "banana": 10, "cherry": 15}
	filtered2 := complexFilter(m22, func(v int) bool { return v > 5 }, func(v int) bool { return v < 15 })
	fmt.Println(filtered2)

	m23 := map[string]int{"a": 1, "b": 2}
	m24 := map[string]int{"b": 3, "c": 4}
	m25 := map[string]int{"a": 5}
	result := mergeAndAggregate(m23, m24, m25)
	fmt.Println(result)

	m26 := map[string]int{"apple": 1, "banana": 2}
	iterateMap(m26, func(key string, value int) {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	})

	people := make(map[string]*Person)
	people["alice"] = &Person{Name: "Alice", Age: 30}
	people["bob"] = &Person{Name: "Bob", Age: 25}
	for key, person := range people {
		fmt.Printf("Key: %s, Person: %+v\n", key, person)
	}

	m27 := map[string]int{"apple": 1, "banana": 2}
	array := mapToArray(m27)
	fmt.Println("Array:", array)
	newMap := arrayToMap(array)
	fmt.Println("New Map:", newMap)

	sg := StudentGrades{grades: make(map[string]int)}
	sg.AddGrade("Math", 85)
	sg.AddGrade("Science", 90)
	fmt.Println("Average Grade:", sg.AverageGrade())

	storage := Storage{records: make(map[int]Record)}
	storage.AddRecord(1, "Alice")
	storage.AddRecord(2, "Bob")
	if record, exists := storage.GetRecord(1); exists {
		fmt.Println("Record:", record)
	}

	cache2 := NewLimitedCache(2)
	cache2.Set("a", "1")
	cache2.Set("b", "2")
	cache2.Set("c", "3") // "a" will be removed
	if value, exists := cache2.Get("a"); exists {
		fmt.Println("a:", value)
	} else {
		fmt.Println("a not found")
	}
	if value, exists := cache2.Get("b"); exists {
		fmt.Println("b:", value)
	}

	sm := StringMap{data: make(map[string][]string)}
	sm.Add("fruits", []string{"apple", "banana"})
	sm.Add("vegetables", []string{"carrot", "pea"})
	fmt.Println("Fruits:", sm.Get("fruits"))
	fmt.Println("Vegetables:", sm.Get("vegetables"))
	sm.Remove("fruits")
	fmt.Println("Fruits after removal:", sm.Get("fruits"))
}
