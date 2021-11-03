package main

import (
	"demo17/employee"
)

func main() {
	e := employee.Init(
		"Chanchai",
		"Ditthapan",
		50,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Lek",
		"Ditthapan",
		50,
		20)

	e.LeavesRemaining()

}

//go mod init demo15
