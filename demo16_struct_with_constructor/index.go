package main

import (
	"demo16/employee"
)

func main() {
	e := employee.New(
		"Chanchai",
		"Ditthapan",
		50,
		20)

	e.LeavesRemaining()

}

//go mod init demo15
