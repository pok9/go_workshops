package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1 somthing")
	}

}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2 somthing")
	}
}

func run3() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run3 somthing")
	}
}

func main() {
	go run1()
	go run2()
	go run3()

	time.Sleep(5 * time.Second)
}
