package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4  <-----

	println()

	newInts := intSeq()
	fmt.Println(newInts()) // 1
	fmt.Println(newInts()) // 2

	println()

	fmt.Println(nextInt()) // 5  <-----

	println("===============================")

	seqString := stringSeq()
	fmt.Println(seqString()) // Y = 1
	fmt.Println(seqString()) // Y = 2

	println("===============================")

	fmt.Println(stringSeq()()) // Y = 1
	fmt.Println(stringSeq()()) // Y = 1
	fmt.Println(stringSeq()()) // Y = 1
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("Y = %d", y)
	}
}
