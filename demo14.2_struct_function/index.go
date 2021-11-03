package main

import "fmt"

func main() {
	var p1 product
	p1.name = "Arduino"
	p1.price = 100
	p1.stock = 20

	show(p1)
	update(&p1)
	show(p1)

	// p1 = p1.clear()
	// testClear(&p1)

	p1 = p1.setDiscount(50)
	show(p1)
}

type product struct {
	name  string
	price int
	stock int
}

//เพิ่ม function ให้ struct แบบไม่มี paremeter
func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

//ทำแบบ pointer
func testClear(p *product) {
	p.price = 0
	p.stock = 0
}

//เพิ่ม function ให้ struct แบบมี paremeter
func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

func show(p product) {
	fmt.Println(p)
}

func update(p *product) { //ทำแบบนี้จะสามารถอัพเดทได้โดยไม่ต้อง return ค่ากลับไป
	p.price = p.price + 100
	p.stock = 10
}
