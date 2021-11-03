package main

import "fmt"

//Global variable
var count int = 0

func main() {
	fmt.Println("Begin")

	// Explicit Declaration ระบุ type อย่างชัดเช่น
	// var ไม่สามารถเปลี่ยนแปลงค่า type ได้
	var tmp1 int = 0
	var tmp2 string = "hello"
	var tmp3 bool = true

	// const เปลี่ยนแปลงค่าไม่ได้
	const tmp4 int = 0

	// Implicit Declration
	//var tmp5 int = 0 เช็ค type ให้เอง ประกาศเป็น type อะไรเริ่มก็จะเป็น type นั้นตลอด
	tmp5 := 0
	tmp6 := "pok555"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)

	fmt.Println(tmp5)
	fmt.Println(tmp6)

}

func run() {
	fmt.Println(count)
	count++
	fmt.Println(count)
}
