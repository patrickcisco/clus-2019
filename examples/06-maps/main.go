package main

import "fmt"

func main() {

	stringMap := map[string]string{}
	stringMap["hello"] = "world"

	for k, v := range stringMap {
		fmt.Println(fmt.Sprintf("key=%v value=%v", k, v))
	}

	interfaceMap := map[string]interface{}{
		"hello":        "world",
		"secretNumber": 87912321,
		"another map": map[string]int{
			"a key":       2,
			"another key": 3,
		},
	}
	for k, v := range interfaceMap {
		fmt.Println(fmt.Sprintf("key=%v value=%v", k, v))
	}
}
