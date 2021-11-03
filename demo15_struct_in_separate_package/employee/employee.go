package employee

import "fmt"

type Employee struct { // -E- ตัวใหญ่เพื่อทำให้ภายนอกมองเห็นได้
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

//สร้าง function ให้กับ struct
func (e Employee) LeavesRemaining() { // -L- ตัวใหญ่เพื่อทำให้ภายนอกมองเห็นได้
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
