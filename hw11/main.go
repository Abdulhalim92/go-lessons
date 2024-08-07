package main

import "fmt"

func main() {
	fmt.Println(concatenateStrings("Hello, ", "World!")) // "Hello, World!"

	fmt.Println(stringLength("Hello, World!")) // 13

	fmt.Println(containsSubstring("Hello, World!", "World")) // true

	fmt.Println(indexOfSubstring("Hello, World!", "World")) // 7

	fmt.Println(replaceSubstring("Hello, World!", "World", "Go")) // "Hello, Go!"

	fmt.Println(trimSpaces("  Hello, World!  ")) // "Hello, World!"

	fmt.Println(toUpper("Hello, World!")) // "HELLO, WORLD!"
	fmt.Println(toLower("Hello, World!")) // "hello, world!"

	fmt.Println(repeatString("Go", 3)) // "GoGoGo"

	fmt.Println(splitString("a,b,c", ",")) // ["a", "b", "c"]

	fmt.Println(equalFold("Go", "go")) // true

	fmt.Println(replaceFirst("Hello, World! World!", "World", "Go")) // "Hello, Go! World!"

	fmt.Println(reverseString("Hello, World!")) // "!dlroW ,olleH"

	fmt.Println(countCharacters("Hello, World!")) // map[H:1 e:1 l:3 o:2 ,:1  :1 W:1 r:1 d:1 !:1]

	fmt.Println(removeCharacter("Hello, World!", 'o')) // "Hell, Wrld!"

	fmt.Println(countWords("Hello, World! This is Go.")) // 5

	fmt.Println(hasPrefix("Hello, World!", "Hello"))  // true
	fmt.Println(hasSuffix("Hello, World!", "World!")) // true

	fmt.Println(removeDuplicates("Hello, World!")) // "Helo, Wrd!"

	fmt.Println(escapeHTML("<div>Hello, World!</div>")) // "&lt;div&gt;Hello, World!&lt;/div&gt;"

	fmt.Println(createAlias("Hello, World! This is Go.")) // "Hello_World_This_is_Go_"

	fmt.Println(sumNumbers("1 2 3 4 5")) // 15

	fmt.Println(reverseWords("Hello, World! This is Go.")) // "Go. is This World! Hello,"

	fmt.Println(countUniqueWords("Hello, World! Hello, Go.")) // 4

	fmt.Println(isPalindrome("A man a plan a canal Panama")) // true

	fmt.Println(shuffleWords("Hello, World! This is Go.")) // Пример: "This Hello, Go. World! is"

	fmt.Println(sortByLength("Hello, World! This is Go.")) // "Go. is This Hello, World!"

	fmt.Println(hashString("Hello, World!")) // Пример: "315f7c6757a4bd193567dbe8209f5a3623f6e8e8d55b344b76cbeb38ff9dd651"

	fmt.Println(allSubstrings("abc")) // ["a", "ab", "abc", "b", "bc", "c"]

	fmt.Println(findAllAnagrams("cbaebabacd", "abc")) // ["cba", "bac"]

	fmt.Println(countWordsAndCharacters("Hello, World! This is Go.")) // 5, 24
}

func extra() {
	sp := &StringProcessor{Content: "   Hello, World!   "}
	fmt.Println(sp.Process(toUpper))    // "   HELLO, WORLD!   "
	fmt.Println(sp.Process(trimSpaces)) // "Hello, World!"

	prefix := "Hello, "
	suffix := "!"
	addPrefix := func(s string) string {
		return prefix + s
	}
	addSuffix := func(s string) string {
		return s + suffix
	}
	f1 := NewFormatter("World", addPrefix)
	f2 := NewFormatter("World", addSuffix)
	fmt.Println(f1.Content) // "Hello, World"
	fmt.Println(f2.Content) // "World!"

	tp := &TextProcessor{Content: "Hello, World! This is Go."}
	fmt.Println(tp.Process(countWords2))      // 5
	fmt.Println(tp.Process(countCharacters2)) // 24

	sm := &StringManipulator{Content: "Hello World"}
	fmt.Println(sm.Manipulate(reverseString)) // "dlroW olleH"
	fmt.Println(sm.Manipulate(replaceSpaces)) // "Hello-World"

	sf := &StringFilter{Content: "Hello, Go World"}
	fmt.Println(sf.Filter(func(s string) bool { return containsWord(s, "Go") })) // true
	fmt.Println(sf.Filter(func(s string) bool { return isLongerThan(s, 10) }))   // true

	te := &TextEditor{Content: "Hello, World"}
	te.Edit(addPrefix)
	fmt.Println(te.Content) // "Prefix: Hello, World"
	te.Edit(addSuffix)
	fmt.Println(te.Content) // "Prefix: Hello, World :Suffix"

	st := &Statistics{Content: "Hello, World! This is Go."}
	fmt.Println(st.Compute(countSpaces))           // 5
	fmt.Println(st.Compute(countUniqueCharacters)) // 16

	r := &Replacer{Content: "Hello 123 World"}
	r.Replace(func(s string) string { return replaceSubstring(s, "World", "Go") })
	fmt.Println(r.Content) // "Hello 123 Go"
	r.Replace(removeDigits)
	fmt.Println(r.Content) // "Hello  World"

	tf := &TextFormatter{Content: "Hello World"}
	tf.Format(underline)
	fmt.Println(tf.Content) // "H_e_l_l_o_ _W_o_r_l_d"
	tf.Format(replaceSpacesWithUnderscore)
	fmt.Println(tf.Content) // "Hello_World"

	sa := &StringAnalyzer{Content: "Hello, Go World!"}
	fmt.Println(sa.Analyze(countVowels))     // 5
	fmt.Println(sa.Analyze(countConsonants)) // 7
}
