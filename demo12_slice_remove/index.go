package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	showSlice(numbers)

	// leading remove
	numbers = numbers[1:]
	showSlice(numbers)
	numbers = numbers[1:]
	showSlice(numbers)

	// trailing remove
	numbers = numbers[0 : len(numbers)-1]
	showSlice(numbers)

	// remove specific index
	numbers = removeIndex(numbers, 2)
	showSlice(numbers)
	numbers = removeIndex(numbers, 1)
	showSlice(numbers)

}

func showSlice(x []int) {
	fmt.Printf("len=%d cap%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
