package main

import "fmt"

func main() {
	//map 1D
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers["one"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "00f"
	fmt.Println(colors)
	fmt.Println(colors["green"])

	//map 2D
	var courses = make(map[string]map[string]int)

	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200
	courses["Android"]["price"] = 100
	courses["Android"]["code"] = 1234

	courses["IOS"] = make(map[string]int)
	courses["IOS"]["price"] = 100
	courses["IOS"]["code"] = 1234

	fmt.Println(courses["IOS"]["price"])
}
