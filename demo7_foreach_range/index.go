package main

import "fmt"

func main() {
	courses := []string{"Android", "ios", "React"}

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	for _, v := range courses {
		fmt.Println(v)

	}
}
