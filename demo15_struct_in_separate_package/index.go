package main

import (
	"demo15/employee"
)

func main() {
	e := employee.Employee{
		FirstName:   "Chanchai",
		LastName:    "Ditthapan",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()

}

//go mod init demo15
