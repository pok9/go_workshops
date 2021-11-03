package main

import "fmt"

func main() {

	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{1, 2, 3, 4}
	var array3 [3]string

	for index, item := range array1 {
		fmt.Println(item, ",", index)
	}
	fmt.Println()
	for _, item := range array2 {
		fmt.Println(item, ",")
	}
	fmt.Println()
	array3[0], array3[1], array3[2] = "android", "ios", "react"
	for _, item := range array3 {
		fmt.Println(item)
	}
}
