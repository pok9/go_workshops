package main

import "fmt"

func main() {

	defer fmt.Println("Finish1") //ทำงานท้ายสุด เช่น เคลีย memory

	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}

}
