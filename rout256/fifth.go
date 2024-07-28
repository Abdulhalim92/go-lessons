package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// JSON prettify
func fifth() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	results := make([]interface{}, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanf(reader, "%d\n", &n)

		var jsonStr string
		for j := 0; j < n; j++ {
			line, _ := reader.ReadString('\n')
			jsonStr += line
		}

		var data interface{}
		json.Unmarshal([]byte(jsonStr), &data)
		data = prettify(data)
		results[i] = data
	}

	output, _ := json.Marshal(results)
	fmt.Fprintln(writer, string(output))
}

func prettify(data interface{}) interface{} {
	switch v := data.(type) {
	case []interface{}:
		var newArray []interface{}
		for _, item := range v {
			item = prettify(item)
			if !isEmpty(item) {
				newArray = append(newArray, item)
			}
		}
		return newArray
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range v {
			value = prettify(value)
			if !isEmpty(value) {
				newMap[key] = value
			}
		}
		return newMap
	default:
		return data
	}
}

func isEmpty(data interface{}) bool {
	switch v := data.(type) {
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	default:
		return false
	}
}
