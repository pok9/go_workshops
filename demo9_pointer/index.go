package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg //เก็บ Address ขอ msg

	fmt.Println(msg)
	fmt.Println(&msg)        //แสดงค่า Address ของ msg
	fmt.Println(*msgPointer) //แสดง value ของ msgPointer
	fmt.Println(msgPointer)  //แสดง Address ของ msgPointer ที่เป็น Address ของ msg ด้วย

	changeMessage(&msg, "new message 1")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 2")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 3")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
