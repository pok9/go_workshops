package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1) //จองไว้ 1 ตำแหน่ง
	ch <- 1                 //send ส่งเข้า ทำให้ chanel เต็ม เพราะจอง buffer ไว้แค่ 1
	fmt.Println("step1")
	fmt.Println(<-ch) // exit ส่งออก ทำให้ chanel ว่าง1 ปล.ถ้าไม่ส่งค่าออกตรงนี้ บรรทัดที่ 14 (ch <- 2) จะเกิด DeadLock

	ch <- 2 //send
	fmt.Println("step2")
	fmt.Println(<-ch) // msg := <-ch

	time.Sleep(1 * time.Second)
}
