package employee

import "fmt"

type employee struct { // -E- ตัวใหญ่เพื่อทำให้ภายนอกมองเห็นได้
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(
	firstname string,
	lastname string,
	totalLeaves int,
	leacesTaken int) employee {
	e := employee{
		FirstName:   firstname,
		LastName:    lastname,
		TotalLeaves: totalLeaves,
		LeavesTaken: leacesTaken,
	}

	return e

}

//สร้าง function ให้กับ struct
func (e employee) LeavesRemaining() { // -L- ตัวใหญ่เพื่อทำให้ภายนอกมองเห็นได้
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
